package strategy

import "iot/message"

type Strategy interface {
	Input([]byte) (message.Message, error)
	Output(message.Message, string)
	Decode([]byte) message.Message
	GetCode() string
	GetDeviceId(data []byte) string
}
