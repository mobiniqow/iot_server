package main

import (
	"bufio"
	"github.com/go-kit/kit/log"
	"iot/brodcaster"
	"iot/device"
	"iot/message"
	"iot/message_broker/gateway"
	"iot/message_broker/rabbitmq"
	"iot/message_broker/strategy"
	"iot/middlerware"
	"iot/middlerware/try_job"
	"iot/server"
	"os"
	"time"
)

const (
	SCHEDULE = "SD"
	SETTINGS = "CD"
)

func main() {
	PORT := 8080
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger, "service", "url", "iot_server", log.DefaultTimestampUTC, "caller", log.DefaultCaller)
	}
	_middleware := middlerware.GetMiddlewareInstance()
	deviceManager := device.Manager{Devices: make([]device.Device, 0), Logger: logger}
	broadCaster := brodcaster.BroadCaster{MiddleWares: _middleware}
	_middleware.Add(&try_job.TryJob{
		TryNumber:     30,
		SleepTime:     5 * time.Second,
		Jobs:          make(map[string]try_job.Job),
		Logger:        logger,
		DeviceManager: &deviceManager,
		BroadCaster:   &broadCaster,
	})

	_gateway := gateway.NewGateway(logger)
	scheduleStrategy := strategy.ScheduleStrategy{
		StrategyCode: SCHEDULE, DeviceManager: &deviceManager}
	settingsStrategy := strategy.SettingsStrategy{
		StrategyCode: SETTINGS, DeviceManager: &deviceManager,
	}
	_gateway.AddStrategy(&scheduleStrategy)
	_gateway.AddStrategy(&settingsStrategy)

	messageBroker := rabbitmq.NewMessageBroker("amqp://guest:guest@localhost", logger, _gateway, &deviceManager, &broadCaster)
	tcpServer := server.New(PORT, logger, messageBroker, &deviceManager, &broadCaster)

	go func() {
		for {
			reader := bufio.NewReader(os.Stdin)
			data, _ := reader.ReadString('\n')
			decoder := message.Decoder{Logger: logger}
			_type, payload, datetime, err := decoder.Decoder([]byte(data))
			if err != nil {

			}
			message := message.Message{
				Extentions: make([]message.Extention, 0),
				Type:       _type,
				Payload:    payload,
				Date:       datetime,
			}
			broadCaster.SendMessage(deviceManager.Devices[0], &message)
		}
	}()

	tcpServer.Run()
}
