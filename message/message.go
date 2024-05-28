package message

type Type []byte

var (
	GET_ID Type = []byte("VV") // 5656
	JOBS   Type = []byte("CD") // 4344
)

// this extentions for middlewares added data to Main Message payloads

type Message struct {
	Date    string
	Payload []byte
	Type    Type
	//Device     device.Device
	Extentions []Extention
}
