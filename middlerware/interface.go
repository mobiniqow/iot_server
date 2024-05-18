package middlerware

import "net"

type Middleware interface {
	Controller()
	Output(net.Conn, error, string) (net.Conn, error)
	Input(net.Conn, error, string) (net.Conn, error)
}
