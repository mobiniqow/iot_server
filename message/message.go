package message

import "iot/device"

type Type string

const (
	GET_ID Type = "VV"
)

type Message struct {
	Date    string
	Payload string
	Type    Type
	Device  device.Device
}
