package brodcaster

import (
	"fmt"
	"iot/device"
	"iot/message"
	"iot/middlerware"
	"iot/utils"
)

type BroadCaster struct {
	MiddleWares *middlerware.Middlewares
}

func (c *BroadCaster) SendMessage(device device.Device, _message *message.Message) error {
	fmt.Printf("_message.Type %s ,_message.Payload %s\n", _message.Type, _message.Payload)
	_, err := c.MiddleWares.Output(device.Conn, _message)
	if err != nil {
		return err
	}
	content := utils.ContentMaker(*_message)
	device.Conn.Write([]byte(content))
	return nil
}
