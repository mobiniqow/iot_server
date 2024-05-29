package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"iot/message"
	"net"
)

func JobKeyGenerator(conn net.Conn, message message.Message) string {
	return fmt.Sprintf("%s:%s%s%s", conn.RemoteAddr().String(), message.Type, message.Payload, message.Date)
}

func ContentMaker(message message.Message) string {
	extentions := make([]byte, 0)
	for _, extention := range message.Extentions {
		extentions = append(extentions, extention.Code[:]...)
	}
	return fmt.Sprintf("%X%X%X%04X\r\n", message.Type, message.Payload, message.Date, extentions)
}

func ByteArrayToInt(byteSlice []byte) (int, error) {
	return int(binary.BigEndian.Uint64(byteSlice)), nil
}

func StringToMap(data string) map[string]interface{} {
	sec := map[string]interface{}{}
	if err := json.Unmarshal([]byte(string(data)), &sec); err != nil {
		panic(err)
	}
	return sec
}
