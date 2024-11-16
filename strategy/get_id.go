package strategy

import (
	"iot/brodcaster"
	"iot/device"
	"iot/message"
	"iot/utils"
)

var (
	ONLINE  = []byte("1")
	OFFLINE = []byte("0")
)

type GetIdStrategy struct {
	StrategyCode  string
	DeviceManager *device.Manager
	BroadCaster   *brodcaster.BroadCaster
}

func (c *GetIdStrategy) MessageBroker(_message string) (message.Message, error) {
	result := c.Decode(_message)
	return result, nil
}

func (c *GetIdStrategy) ClientHandler(data string) (message.Message, error) {
	//device, _ := c.DeviceManager.GetDeviceByDeviceId(deviceId)
	//c.BroadCaster.SendMessage(device, &_message)

	_type := data[:2]
	payload := data[2:]
	return message.Message{
		Type:       _type,
		Payload:    payload,
		Date:       "",
		Extentions: make([]message.Extention, 0),
	}, nil
}

func (c *GetIdStrategy) GetCode() string {
	return c.StrategyCode
}

// device id ro hamrah ba message bar migardone
func (c *GetIdStrategy) Decode(data string) message.Message {
	dataString := string(data)
	dataMap := utils.StringToMap(dataString)
	_type := (dataMap["type"].(string))
	payload := dataMap["payload"].(string)
	//byteOfPayload, _ := utils.HexToByte(payload)
	_message := message.NewMessage(_type, "", payload)
	return *_message
	// dar payload 4 caracter aval shomare relay hast va baghie barname haftegi relay
}

func (c *GetIdStrategy) GetDeviceId(data string) string {
	dataString := string(data)
	dataMap := utils.StringToMap(dataString)
	deviceId := dataMap["device_id"].(string)
	return deviceId
	// dar payload 4 caracter aval shomare relay hast va baghie barname haftegi relay
}
