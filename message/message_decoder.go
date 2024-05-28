package message

import (
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
func (dec *Decoder) Decoder(bytes []byte) (string, string) {
	dec.Logger.Log("Decoder bytes:", bytes)
	_type := string(bytes[:2])
	payload := string(bytes[2:])
	return _type, payload
}

func SplitMessage(data string) (Type, []byte, error) {
	num := Type(data[:2])
	return num, []byte(data[2:]), nil
}
