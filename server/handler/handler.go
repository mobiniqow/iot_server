package handler

import (
	"fmt"
	"iot/device"
	"iot/message"
	"iot/message_broker/gateway"
	"iot/message_broker/rabbitmq"
	"iot/middlerware"
	"iot/strategy"
	"net"
	"time"

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
	Gateway       *gateway.Gateway
}

func (c *Handler) CloseConnection(Connection net.Conn) {
	err := Connection.Close()
	if err != nil {
		c.Logger.Log("Connection closed with error: %v", err)
		print(err)
	} else {
		c.Logger.Log("Connection closed")
	}
	c.DeviceManager.Delete(c.Device)
}
func (c *Handler) DisconnectOldDevice(device device.Device) {
	// ابتدا دستگاه قدیمی را پیدا می‌کنیم
	oldDevice, err := c.DeviceManager.GetDeviceByDeviceId(string(device.DeviceID))
	if err == nil {
		// دستگاه قدیمی پیدا شد، آن را حذف می‌کنیم
		c.DeviceManager.Delete(oldDevice)
		c.Logger.Log("Disconnected old device with DeviceID:", oldDevice.DeviceID)
	}
}

func (c *Handler) Start() {
	go func() {
		defer c.CloseConnection(c.Connection)
		go func() {
			sendMessage := []byte("RG\r\n")
			c.Device.Conn.Write(sendMessage)
			time.Sleep(10 * time.Second)
			device, _ := c.DeviceManager.GetDeviceByConnection(c.Connection)
			if !device.IsValid() {
				fmt.Print("raft to va valid bood")
				c.Device.Conn.Write(sendMessage)
			}
			time.Sleep(10 * time.Second)
			device, _ = c.DeviceManager.GetDeviceByConnection(c.Connection)
			if !device.IsValid() {
				c.Device.Conn.Write(sendMessage)
			}
		}()
		for {
			buffer := make([]byte, 1024)
			n, err := c.Connection.Read(buffer)
			if err != nil {
				c.Logger.Log("error from reading data %s from socket: %v", err, c.Connection)
				return
			} else {
				body := buffer[:n]
				_message, err := c.Gateway.ClientHandler(string(body))
				if err != nil {
					c.Logger.Log("error from reading data %s from socket: %v", err, c.Connection)
				} else {
					if _message.Type == strategy.GET_ID {
						// به‌روزرسانی DeviceID
						c.Device.DeviceID = []byte(_message.Payload)

						// قطع ارتباط و حذف دستگاه قبلی
						c.DisconnectOldDevice(c.Device)

						// دستگاه جدید را اضافه کنید
						c.DeviceManager.Update(c.Device)

						c.Logger.Log("Received message with new DeviceID:", c.Device.DeviceID)
						c.MessageBroker.SendData(string(c.Device.DeviceID), _message)
					} else {
						// اگر پیام نوع دیگری داشت
						fmt.Printf("\n messageinputelse %x %s \n", _message.Payload, _message.Payload)
						c.MessageBroker.SendData(string(c.Device.DeviceID), _message)
					}
				}
			}
		}
	}()
}
