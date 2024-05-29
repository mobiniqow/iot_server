package message

import (
	"errors"
	"fmt"
	"github.com/go-kit/kit/log"
)

type Decoder struct {
	Logger log.Logger
}

func (dec *Decoder) Encoder(message Message) []byte {
	dec.Logger.Log("Encode message:", message)
	_type := message.Type
	payload := message.Payload
	date := message.Date
	template := fmt.Sprintf("%x%x%x", _type, payload, date)
	return []byte(template)
}

//func (dec *Decoder) Decoder(bytes []byte) (string, string, error) {
//	dec.Logger.Log("Decoder bytes:", bytes)
//
//	_type := string(bytes[:2])
//	payload := string(bytes[2:])
//	return _type, payload, nil
//}

func (dec *Decoder) Decoder(data []byte) (Type, []byte, error) {
	if len(data) < 2 {
		return Type{}, nil, errors.New("message too short")
	}
	num := Type(data[:2])
	return num, []byte(data[2:]), nil
}
