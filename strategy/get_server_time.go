package strategy

import (
	"fmt"
	"iot/brodcaster"
	"iot/device"
	"iot/message"
	"iot/utils"
)

type GetServerTimeStrategy struct {
	StrategyCode  string
	DeviceManager *device.Manager
	BroadCaster   *brodcaster.BroadCaster
}

func (c *GetServerTimeStrategy) MessageBroker(_message []byte) (message.Message, error) {
	result := c.Decode(_message)
	return result, nil
}

func (c *GetServerTimeStrategy) ClientHandler(data []byte) (message.Message, error) {
	_type := data[:2]
	payload := data[2:]
	return message.Message{
		Type:       _type,
		Payload:    payload,
		Date:       nil,
		Extentions: make([]message.Extention, 0),
	}, nil
}

func (c *GetServerTimeStrategy) GetCode() string {
	return c.StrategyCode
}

// device id ro hamrah ba message bar migardone
func (c *GetServerTimeStrategy) Decode(data []byte) message.Message {
	println("GetServerTimeStrategy")
	dataString := string(data)
	dataMap := utils.StringToMap(dataString)
	_type := []byte(dataMap["type"].(string))
	payload := dataMap["payload"].(string)
	fmt.Printf("\npayload %v \r\n", payload)
	byteOfPayload := []byte(payload)
	_message := message.NewMessage(_type, nil, byteOfPayload)
	return *_message
	// dar payload 4 caracter aval shomare relay hast va baghie barname haftegi relay
}

func (c *GetServerTimeStrategy) GetDeviceId(data []byte) string {
	println("GetServerTimeStrategy")
	dataString := string(data)
	dataMap := utils.StringToMap(dataString)
	deviceId := dataMap["device_id"].(string)
	return deviceId
	// dar payload 4 caracter aval shomare relay hast va baghie barname haftegi relay
}
