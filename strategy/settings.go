package strategy

import (
	"fmt"
	"iot/brodcaster"
	"iot/device"
	"iot/message"
	"iot/utils"
)

type SettingsStrategy struct {
	StrategyCode  string
	DeviceManager *device.Manager
	BroadCaster   *brodcaster.BroadCaster
}

func (c *SettingsStrategy) MessageBroker(_message []byte) (message.Message, error) {
	result := c.Decode(_message)
	return result, nil
}

func (c *SettingsStrategy) ClientHandler(data []byte) (message.Message, error) {
	//device, _ := c.DeviceManager.GetDeviceByDeviceId(deviceId)
	//c.BroadCaster.SendMessage(device, &_message)
	fmt.Printf("stirngdata %s\n", data)
	_type := data[:2]
	payload := make([]byte, 0)
	datetime := make([]byte, 0)

	if len(data) >= 4 {
		payload = data[2:4]
		datetime = data[4:]
	}

	return message.Message{
		Payload:    payload,
		Type:       _type,
		Date:       datetime,
		Extentions: make([]message.Extention, 0),
	}, nil
}

func (c *SettingsStrategy) GetCode() string {
	return c.StrategyCode
}

// device id ro hamrah ba message bar migardone
func (c *SettingsStrategy) Decode(data []byte) message.Message {
	dataString := string(data)
	dataMap := utils.StringToMap(dataString)
	_type := []byte(dataMap["type"].(string))
	datetime := []byte(dataMap["datetime"].(string))
	payload := dataMap["payload"].(string)
	byteOfPayload, _ := utils.HexToByte(payload)
	_message := message.NewMessage(_type, datetime, byteOfPayload)
	return *_message
	// dar payload 4 caracter aval shomare relay hast va baghie barname haftegi relay
}

func (c *SettingsStrategy) GetDeviceId(data []byte) string {
	dataString := string(data)
	dataMap := utils.StringToMap(dataString)
	deviceId := dataMap["device_id"].(string)
	return deviceId
	// dar payload 4 caracter aval shomare relay hast va baghie barname haftegi relay
}
