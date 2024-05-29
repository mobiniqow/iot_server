package handler

import (
	"bytes"
	"iot/device"
	"iot/message"
	"iot/message_broker/rabbitmq"
	"iot/middlerware"
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
	Middleware    *middlerware.Middlewares
	MessageBroker rabbitmq.MessageBroker
}

// handle close connection device
func (c *Handler) CloseConnection(Connection net.Conn) {
	err := Connection.Close()
	if err != nil {
		c.Logger.Log("Connection closed with error: %v", err)
		panic(err)
	} else {
		c.Logger.Log("Connection closed")
	}
	c.DeviceManager.Delete(c.Device)
}

func (h *Handler) Start() {
	go func() {
		defer h.CloseConnection(h.Connection)
		// buffer for reading data from socket
		buffer := make([]byte, 1024)

		for {
			n, err := h.Connection.Read(buffer)
			if err != nil {
				h.Logger.Log("error from reading data %s from socket: %v", err, h.Connection)
			} else {
				body := buffer[:n]
				_type, payload, err := h.Decoder.Decoder(body)
				msg := message.NewMessage(_type, payload)
				if h.Validator.Validate(body) {
					if err != nil {
						h.Logger.Log("error from reading data %s from socket: %v", err, h.Connection)
					} else {
						// agar payami ke amad moshabeh id bod ono add konam be device haie mojod
						if bytes.Equal(_type, message.GET_ID) {
							h.Device.DeviceID = payload
							h.DeviceManager.Update(h.Device)
							h.Logger.Log("Received message with ID: %v", h.Device.DeviceID)
						} else {
							if h.Device.IsValid() {
								h.MessageBroker.SendData(string(h.Device.DeviceID), *msg)
								print("yes bro")
							}
						}
					}
				}

				if err != nil {
					h.Logger.Log("error from reading data %s from socket: %v", err, h.Connection)
				} else {
					_, err := h.Middleware.Inputs(h.Connection, msg)
					if err != nil {
						h.Logger.Log("Get message with error: %v", err)
						return
					}
				}
			}
		}
	}()
}
