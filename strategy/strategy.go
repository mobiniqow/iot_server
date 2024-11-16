package strategy

import "iot/message"

const (
	SCHEDULE    = "SD" // 5344
	SETTINGS    = "CD" // 4344
	GET_ID      = "VV" // 5656
	JOBS        = "JS" // 4a53
	SERVER_TIME = "ST" // 4a53
)

type Strategy interface {
	MessageBroker(string) (message.Message, error)
	ClientHandler(string) (message.Message, error)
	Decode(string) message.Message
	GetCode() string
	GetDeviceId(string) string
}
