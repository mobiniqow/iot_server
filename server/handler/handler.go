package handler

import (
	"fmt"
	"iot/device"
	"iot/message"
	"iot/server"
	"net"

	"github.com/go-kit/kit/log"
)

type Handler struct {
	Connection    net.Conn
	Logger        log.Logger
	DeviceManager *device.Manager
	Device        device.Device
	Validator     message.Validator
	Decoder       message.Decoder
}

func (h *Handler) Start() {
	middlewares := server.GetMiddlewareInstance().Middleware
	go func() {
		defer func(Connection net.Conn) {
			err := Connection.Close()
			if err != nil {
				h.Logger.Log("Connection closed with error: %v", err)
				panic(err)
			} else {
				h.Logger.Log("Connection closed")
			}
			h.DeviceManager.Delete(h.Device)
		}(h.Connection)
		buffer := make([]byte, 1024)
		for {
			n, err := h.Connection.Read(buffer)
			if err != nil {
				fmt.Println("Error:", err)
				h.Logger.Log("Get message with error: %v", err)
				return
			}
			body := buffer[:n]
			fmt.Printf("Received: %v\n", body)
			if h.Validator.Validate(body) {
				status, payload := h.Decoder.Decoder(body)
				if status == string(message.GET_ID) {
					h.Device.DeviceID = payload
					h.Logger.Log("Received message with ID: %v", h.Device.DeviceID)
					h.DeviceManager.Update(h.Device)
				} else {
					if h.Device.IsValid() {
						print("yes bro")
					}
				}
			}
		}
	}()
}
