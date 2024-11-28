package rabbitmq

import (
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/wagslane/go-rabbitmq"
	"iot/brodcaster"
	"iot/device"
	"iot/message_broker/gateway"
	"iot/utils"
)

type ConsumerHandler struct {
	Logger        log.Logger
	DeviceManager *device.Manager
	Gateway       *gateway.Gateway
	BroadCaster   *brodcaster.BroadCaster
}

func (c *ConsumerHandler) Handler(d rabbitmq.Delivery) rabbitmq.Action {
	c.Logger.Log("consumed: %v", string(d.Body))
	sec := utils.StringToMap(string(d.Body))
	deviceId := sec["device_id"].(string)
	_message, _ := c.Gateway.MessageBroker(string(d.Body))
	//fmt.Printf("date %s \n", date)'
	fmt.Printf("_message %v\n", _message)
	device, err := c.DeviceManager.GetDeviceByDeviceId(deviceId)
	if err != nil {
		println(err.Error())
	} else {
		c.BroadCaster.SendMessage(device, &_message)
	}
	fmt.Print(sec)
	return rabbitmq.Ack
}
