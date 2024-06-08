package strategy

import (
	"crypto/md5"
	"fmt"
	"io"
	"iot/brodcaster"
	"iot/device"
	"iot/message"
	"iot/utils"
)

const DATE_TIME_LENGTH = 19

// 20 character baraye payloadi ke daraye relay va timereshe
const PAYLOAD_WITH_DATE_TIME_LENGTH = 20

type ScheduleStrategy struct {
	StrategyCode  string
	DeviceManager *device.Manager
	BroadCaster   *brodcaster.BroadCaster
}

func (c *ScheduleStrategy) MessageBroker(_message []byte) (message.Message, error) {
	result := c.Decode(_message)
	return result, nil
}
func md(str string) string {
	h := md5.New()
	io.WriteString(h, str)

	return fmt.Sprintf("%x", h.Sum(nil))
}
func (c *ScheduleStrategy) ClientHandler(data []byte) (message.Message, error) {
	//device, _ := c.DeviceManager.GetDeviceByDeviceId(deviceId)
	//c.BroadCaster.SendMessage(device, &_message)

	_type := data[:2]
	payload := data[2:]
	datetime := make([]byte, 0)
	if len(data) >= DATE_TIME_LENGTH+PAYLOAD_WITH_DATE_TIME_LENGTH+len(_type) {
		payload = data[2 : len(data)-DATE_TIME_LENGTH]
		datetime = data[len(data)-DATE_TIME_LENGTH:]
	}
	fmt.Printf("datetime:%x\n", datetime)

	md(string(payload))
	fmt.Printf("%08b\n", datetime)
	return message.Message{
		Type:    _type,
		Payload: data,
		Date:    datetime,
	}, nil
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
	byteOfPayload, _ := utils.HexToByte(payload)
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
