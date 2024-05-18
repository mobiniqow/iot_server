package middlerware

import (
	"net"
)

type Middlewares struct {
	Middleware []Middleware
}

func GetMiddlewareInstance() *Middlewares {

	return &Middlewares{
		Middleware: make([]Middleware, 0),
	}
}

func (m *Middlewares) Add(middleware Middleware) {
	m.Middleware = append(m.Middleware, middleware)
}
func (m *Middlewares) Inputs(con net.Conn, data string) (c net.Conn, err error) {
	for _, m2 := range m.Middleware {
		m2.Input(con, err, data)
	}
	return con, err
}
