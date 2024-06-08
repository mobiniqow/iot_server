package message

type Type []byte

// this extentions for middlewares added data to Main Message payloads

type Message struct {
	Payload    []byte
	Type       Type
	Date       []byte
	Extentions []Extention
}

func NewMessage(_type, datetime, payload []byte) *Message {
	return &Message{Payload: payload, Type: _type, Extentions: make([]Extention, 0), Date: datetime}
}
