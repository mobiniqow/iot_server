package main

import (
	"github.com/go-kit/kit/log"
	"iot/brodcaster"
	"iot/device"
	"iot/message_broker/gateway"
	"iot/message_broker/rabbitmq"
	"iot/middlerware"
	//	"iot/middlerware/try_job"
	"iot/server"
	"iot/strategy"
	"os"
	//	"time"
)

func main() {
	PORT := 9090
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger, "service", "url", "iot_server", log.DefaultTimestampUTC, "caller", log.DefaultCaller)
	}
	_middleware := middlerware.GetMiddlewareInstance()
	deviceManager := device.Manager{Devices: make([]device.Device, 0), Logger: logger}
	broadCaster := brodcaster.BroadCaster{MiddleWares: _middleware}
	//	_middleware.Add(&try_job.TryJob{
	//		TryNumber:     4,
	//		SleepTime:     3 * time.Second,
	//		Jobs:          make(map[string]try_job.Job),
	//		Logger:        logger,
	//		DeviceManager: &deviceManager,
	//		BroadCaster:   &broadCaster,
	//	})

	_gateway := gateway.NewGateway(logger)
	scheduleStrategy := strategy.ScheduleStrategy{
		StrategyCode: strategy.SCHEDULE, DeviceManager: &deviceManager}

	settingsStrategy := strategy.SettingsStrategy{
		StrategyCode: strategy.SETTINGS, DeviceManager: &deviceManager,
	}

	getIdStrategy := strategy.GetIdStrategy{
		StrategyCode: strategy.GET_ID, DeviceManager: &deviceManager,
	}

	serverTime := strategy.GetServerTimeStrategy{
		StrategyCode: strategy.SERVER_TIME, DeviceManager: &deviceManager,
	}

	lastState := strategy.GetDeviceLastState{
		StrategyCode: strategy.LAST_STATE, DeviceManager: &deviceManager,
	}

	_gateway.AddStrategy(&getIdStrategy)
	_gateway.AddStrategy(&scheduleStrategy)
	_gateway.AddStrategy(&settingsStrategy)
	_gateway.AddStrategy(&serverTime)
	_gateway.AddStrategy(&lastState)

	messageBroker := rabbitmq.NewMessageBroker("amqp://guest:guest@localhost", logger, _gateway, &deviceManager, &broadCaster)
	tcpServer := server.New(PORT, logger, messageBroker, &deviceManager, &broadCaster, _gateway)

	//go func() {
	//	for {
	//		reader := bufio.NewReader(os.Stdin)
	//		data, _ := reader.ReadString('\n')
	//		message, err := _gateway.ClientHandler([]byte(data))
	//		if err != nil {
	//
	//		}
	//		broadCaster.SendMessage(deviceManager.Devices[0], &message)
	//	}
	//}()

	tcpServer.Run()
}
