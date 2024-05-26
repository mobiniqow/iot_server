package utils

import (
	"fmt"
	"iot/message"
	"net"
)

func JobKeyGenerator(conn net.Conn, message message.Message) string {
	return fmt.Sprintf("%s:%s%s", conn.RemoteAddr().String(), message.Type, message.Payload)
}

func ContentMaker(message message.Message) string {
	extentions := ""
	for _, extention := range message.Extentions {
		extentions += extention.Code
	}
	return fmt.Sprintf("%X%X%X\r\n", message.Type, message.Payload, extentions)
}
