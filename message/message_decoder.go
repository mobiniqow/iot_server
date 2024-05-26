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
	status := message.Type
	payload := message.Payload
	date := message.Date
	template := fmt.Sprintf("%x%x%x", status, payload, date)
	return []byte(template)
}
func (dec *Decoder) Decoder(bytes []byte) (string, string) {
	dec.Logger.Log("Decoder bytes:", bytes)
	status := string(bytes[:2])
	payload := string(bytes[2:])
	return status, payload
}

func SplitMessage(data string) (Type, string, error) {
	num := Type(data[:2])
	fmt.Printf("number %s \n", num)
	return Type(num), data[4:], nil
}
