package main

import (
	"bufio"
	"fmt"
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
	middlerware := middlerware.GetMiddlewareInstance()
	middlerware.Add(&try_job.TryJob{
		TryNumber: 3,
		SleepTime: 1 * time.Second,
		Jobs:      make(map[string]try_job.Job),
	})
	tcpServer := server.New(PORT, *middlerware)
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
