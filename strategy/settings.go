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

func (c *SettingsStrategy) MessageBroker(_message string) (message.Message, error) {
	result := c.Decode(_message)
	return result, nil
}

func (c *SettingsStrategy) ClientHandler(data string) (message.Message, error) {
	//device, _ := c.DeviceManager.GetDeviceByDeviceId(deviceId)
	//c.BroadCaster.SendMessage(device, &_message)
	fmt.Printf("stirngdata %s\n", data[:2])
	_type := data[:2]
	var payload string
	datetime := ""

	if len(data) >= 4 {
		d := fmt.Sprintf("%s", data[2:12])
		fmt.Printf("asdsda %s\n", d)
		payload = (d)
		datetime = data[12:]
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
func (c *SettingsStrategy) Decode(data string) message.Message {
	dataString := string(data)
	dataMap := utils.StringToMap(dataString)
	_type := (dataMap["type"].(string))
	datetime := (dataMap["datetime"].(string))
	payload := dataMap["payload"].(string)
	//byteOfPayload, _ := utils.HexToByte(payload)
	_message := message.NewMessage(_type, datetime, payload)
	return *_message
	// dar payload 4 caracter aval shomare relay hast va baghie barname haftegi relay
}

func (c *SettingsStrategy) GetDeviceId(data string) string {
	dataString := data
	dataMap := utils.StringToMap(dataString)
	deviceId := dataMap["device_id"].(string)
	return deviceId
	// dar payload 4 caracter aval shomare relay hast va baghie barname haftegi relay
}
