package message

type Type []byte

var (
	GET_ID Type = []byte("VV") // 5656
	JOBS   Type = []byte("CD") // 4344
	//ORDER    Type = []byte("CC") // 4343
	//SCHEDULE Type = []byte("CB") // 4342
)

// this extentions for middlewares added data to Main Message payloads

type Message struct {
	Date    string
	Payload []byte
	Type    Type
	//Device     device.Device
	Extentions []Extention
}

func NewMessage(_type, payload []byte) *Message {

	return &Message{Payload: payload, Type: _type, Extentions: make([]Extention, 0)}
}
