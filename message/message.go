package message

type Type string

const (
	GET_ID Type = "VV"
	JOBS        = "CD"
)

// this extentions for middlewares added data to Main Message payloads

type Message struct {
	Date    string
	Payload string
	Type    Type
	//Device     device.Device
	Extentions []Extention
}
