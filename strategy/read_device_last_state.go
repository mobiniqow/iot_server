package strategy

import (
	"fmt"
	"iot/brodcaster"
	"iot/device"
	"iot/message"
	"iot/utils"
)

type GetDeviceLastState struct {
	StrategyCode  string
	DeviceManager *device.Manager
	BroadCaster   *brodcaster.BroadCaster
}

func (c *GetDeviceLastState) MessageBroker(_message string) (message.Message, error) {
	result := c.Decode(_message)
	return result, nil
}

func (c *GetDeviceLastState) ClientHandler(data string) (message.Message, error) {
	_type := data[:2]
	payload := data[2:]
	return message.Message{
		Type:       _type,
		Payload:    payload,
		Date:       "",
		Extentions: make([]message.Extention, 0),
	}, nil
}

func (c *GetDeviceLastState) GetCode() string {
	return c.StrategyCode
}

// device id ro hamrah ba message bar migardone
func (c *GetDeviceLastState) Decode(data string) message.Message {
	println("GetDeviceLastState")
	dataString := string(data)
	dataMap := utils.StringToMap(dataString)
	_type := (dataMap["type"].(string))
	payload := dataMap["payload"].(string)
	fmt.Printf("\npayload %v \r\n", payload)
	byteOfPayload := (payload)
	_message := message.NewMessage(_type, "", byteOfPayload)
	return *_message
	// dar payload 4 caracter aval shomare relay hast va baghie barname haftegi relay
}

func (c *GetDeviceLastState) GetDeviceId(data string) string {
	println("GetDeviceLastState")
	dataString := string(data)
	dataMap := utils.StringToMap(dataString)
	deviceId := dataMap["device_id"].(string)
	return deviceId
	// dar payload 4 caracter aval shomare relay hast va baghie barname haftegi relay
}
