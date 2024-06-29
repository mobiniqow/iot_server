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
	MessageBroker([]byte) (message.Message, error)
	ClientHandler([]byte) (message.Message, error)
	Decode([]byte) message.Message
	GetCode() string
	GetDeviceId(data []byte) string
}
