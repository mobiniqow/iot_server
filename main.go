package main

import (
	"iot/middlerware/try_job"
	"iot/server"
	"time"
)

func main() {
	//a := uint16(65535)
	//
	//h := fmt.Sprintf("%x", a)
	//print(strconv.FormatInt(int64(a), 2))

	//
	PORT := 8080
	tcpServer := server.New(PORT)
	tcpServer.Use(&try_job.TryJob{
		TryNumber: 3,
		SleepTime: 6 * time.Second,
		Jobs:      make(map[string]try_job.Job),
	})
	tcpServer.Run()
}
