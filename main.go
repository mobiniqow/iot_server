package main

import (
	"bufio"
	"fmt"
	"github.com/go-kit/kit/log"
	"iot/device"
	"iot/message"
	"iot/middlerware"
	"iot/middlerware/try_job"
	"iot/server"
	"os"
	"time"
)

func main() {
	PORT := 8080

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger, "service", "url", "iot_server", log.DefaultTimestampUTC, "caller", log.DefaultCaller)
	}
	middlerware := middlerware.GetMiddlewareInstance()

	middlerware.Add(&try_job.TryJob{
		TryNumber: 30,
		SleepTime: 4 * time.Second,
		Jobs:      make(map[string]try_job.Job),
		Logger:    logger,
	})
	tcpServer := server.New(PORT, logger, *middlerware)

	go func() {
		for {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Enter text: ")
			data, _ := reader.ReadString('\n')
			dm := device.GetInstanceManagerWithoutLogger()
			_type, payload, err := message.SplitMessage(data)
			if err != nil {

			}
			message := message.Message{
				Extentions: make([]message.Extention, 0),
				Type:       _type,
				Payload:    payload,
			}
			dm.SendMessage(dm.Devices[0].Conn, &message)
		}
	}()
	tcpServer.Run()
}
