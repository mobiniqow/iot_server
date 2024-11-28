package strategy

import "iot/message"

const (
	SCHEDULE    = "RS"
	SETTINGS    = "RR"
	GET_ID      = "RG"
	JOBS        = "JS"
	SERVER_TIME = "RT"
	LAST_STATE  = "RR"
)

type Strategy interface {
	MessageBroker(string) (message.Message, error)
	ClientHandler(string) (message.Message, error)
	Decode(string) message.Message
	GetCode() string
	GetDeviceId(string) string
}
