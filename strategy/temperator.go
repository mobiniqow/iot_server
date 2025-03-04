package strategy

import (
	"iot/brodcaster"
	"iot/device"
	"iot/message"
	"iot/utils"
)

type TemperatureStrategy struct {
	StrategyCode  string
	DeviceManager *device.Manager
	BroadCaster   *brodcaster.BroadCaster
}

func (c *TemperatureStrategy) MessageBroker(_message string) (message.Message, error) {
	result := c.Decode(_message)
	return result, nil
}

func (c *TemperatureStrategy) ClientHandler(data string) (message.Message, error) {
	_type := data[:2]
	payload := data[2:]
	datetime := ""
	if len(data) == 28 {
		datetime = data[28:]
	}
	return message.Message{
		Type:       _type,
		Payload:    payload,
		Date:       datetime,
		Extentions: make([]message.Extention, 0),
	}, nil
}

func (c *TemperatureStrategy) GetCode() string {
	return c.StrategyCode
}

// device id ro hamrah ba message bar migardone
func (c *TemperatureStrategy) Decode(data string) message.Message {
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

func (c *TemperatureStrategy) GetDeviceId(data string) string {
	dataString := data
	dataMap := utils.StringToMap(dataString)
	deviceId := dataMap["device_id"].(string)
	return deviceId
	// dar payload 4 caracter aval shomare relay hast va baghie barname haftegi relay
}
