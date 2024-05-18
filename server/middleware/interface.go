package middleware

import "net"

type Middleware interface {
	Output(net.Conn, error) (net.Conn, error)
	Input(net.Conn, error) (net.Conn, error)
}
