package rabbitmq

import (
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/wagslane/go-rabbitmq"
	"iot/device"
	"iot/message_broker/gateway"
	"iot/utils"
)

type ConsumerHandler struct {
	Logger        log.Logger
	DeviceManager *device.Manager
	Gateway       *gateway.Gateway
}

func (c *ConsumerHandler) Handler(d rabbitmq.Delivery) rabbitmq.Action {
	c.Logger.Log("consumed: %v", string(d.Body))
	sec := utils.StringToMap(string(d.Body))
	deviceId := sec["device_id"].(string)
	_message, _ := c.Gateway.Input(d.Body)
	//fmt.Printf("date %s \n", date)'

	fmt.Printf("_message %s\n", _message)
	c.DeviceManager.SendMessageWithDeviceId(deviceId, _message)
	return rabbitmq.Ack
}
