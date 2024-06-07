package strategy

import (
	"iot/device"
	"iot/message"
	"iot/utils"
)

type ScheduleStrategy struct {
	StrategyCode  string
	DeviceManager *device.Manager
}

func (c *ScheduleStrategy) Input(_message []byte) (message.Message, error) {
	result := c.Decode(_message)
	return result, nil
}

func (c *ScheduleStrategy) Output(_message message.Message, deviceId string) {
	device, _ := c.DeviceManager.GetDeviceByDeviceId(deviceId)
	c.DeviceManager.SendMessage(device, &_message)
}

func (c *ScheduleStrategy) GetCode() string {
	return c.StrategyCode
}

// device id ro hamrah ba message bar migardone
func (c *ScheduleStrategy) Decode(data []byte) message.Message {
	dataString := string(data)
	dataMap := utils.StringToMap(dataString)
	_type := []byte(dataMap["type"].(string))
	datetime := []byte(dataMap["datetime"].(string))
	payload := dataMap["payload"].(string)
	payloadHex := ""
	DAY_WEEK_COUNT := 7
	for i := 0; i <= DAY_WEEK_COUNT; i++ {
		dayBinary := payload[i : i+DAY_WEEK_COUNT]
		payloadHex += utils.BinaryToHex(dayBinary)
	}
	println(payloadHex)
	byteOfPayload, _ := utils.HexToByte(payloadHex)
	_message := message.NewMessage(_type, datetime, byteOfPayload)
	return *_message
	// dar payload 4 caracter aval shomare relay hast va baghie barname haftegi relay
}
func (c *ScheduleStrategy) GetDeviceId(data []byte) string {
	dataString := string(data)
	dataMap := utils.StringToMap(dataString)
	deviceId := dataMap["device_id"].(string)
	return deviceId
	// dar payload 4 caracter aval shomare relay hast va baghie barname haftegi relay
}
