package main

import (
	"encoding/hex"
	"fmt"
	"github.com/go-kit/kit/log"
	"iot/device"
	"iot/message"
	"iot/server/handler"
	"net"
	"os"
)

type MyEnum string

const (
	Foo MyEnum = "VV"
	Bar MyEnum = "BB"
)

func main() {
	a := "5656"
	bs, err := hex.DecodeString(a)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bs))
	// Listen for incoming connections
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer listener.Close()
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger, "service", "url", "iot_server", log.DefaultTimestampUTC, "caller", log.DefaultCaller)
	}

	fmt.Println("Server is listening on port 8080")

	deviceManager := device.GetInstanceManager(logger)
	validator := message.Validator{}
	decoder := message.Decoder{Logger: logger}
	for {
		// Accept incoming connections
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		newDevice := device.Device{Conn: conn, ClientID: conn.RemoteAddr().String()}
		deviceManager.Add(newDevice)
		handler := handler.Handler{Connection: conn, DeviceManager: deviceManager, Logger: logger, Validator: validator, Decoder: decoder, Device: newDevice}
		handler.Start()
	}
}
