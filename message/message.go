package message

type Type []byte

// this extentions for middlewares added data to Main Message payloads

type Message struct {
	Payload    string
	Type       string
	Date       string
	Extentions []Extention
}

func NewMessage(_type, datetime, payload string) *Message {
	return &Message{Payload: payload, Type: _type, Extentions: make([]Extention, 0), Date: datetime}
}
