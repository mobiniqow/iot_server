package server

import (
	"iot/server/middleware"
	"sync"
)

type Middlewares struct {
	Middleware []middleware.Middleware
}

var ones sync.Mutex
var instance *Middlewares

func GetMiddlewareInstance() *Middlewares {
	ones.Lock()
	defer ones.Unlock()
	if instance == nil {
		instance = &Middlewares{
			Middleware: make([]middleware.Middleware, 0),
		}
	}
	return instance
}

func (m *Middlewares) Add(middleware middleware.Middleware) {
	m.Middleware = append(m.Middleware, middleware)
}
