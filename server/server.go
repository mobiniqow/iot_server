package server

import (
	"fmt"
	"iot/device"
	"iot/message"
	"iot/middlerware"
	"iot/server/handler"
	"net"
	"os"

	"github.com/go-kit/kit/log"
)

type server struct {
	Port        int
	middlewares middlerware.Middlewares
}

func New(port int) *server {
	return &server{
		Port:        port,
		middlewares: *middlerware.GetMiddlewareInstance(),
	}
}

func (s *server) Run() {
	for _, middleware := range s.middlewares.Middleware {
		middleware.Controller()
	}
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger, "service", "url", "iot_server", log.DefaultTimestampUTC, "caller", log.DefaultCaller)
	}
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", s.Port))
	if err != nil {
		logger.Log("Error", err)
		return
	}

	if err != nil {
		logger.Log("Error:", err)
		return
	}
	defer listener.Close()
	defer logger.Log("server shutdown")
	// fmt.Println("Server is listening on port 8080")
	logger.Log("Server is listening on port 8080")

	deviceManager := device.GetInstanceManager(logger)

	validator := message.Validator{}

	decoder := message.Decoder{Logger: logger}

	for {
		// Accept incoming connections
		conn, err := listener.Accept()
		if err != nil {
			logger.Log("Error:", err)
			continue
		}
		newDevice := device.Device{Conn: conn, ClientID: conn.RemoteAddr().String()}
		deviceManager.Add(newDevice)
		handler := handler.Handler{Connection: conn, DeviceManager: deviceManager, Logger: logger, Validator: validator,
			Decoder: decoder, Device: newDevice, Middleware: &s.middlewares}
		handler.Start()

	}
}

func (s *server) Use(middleware middlerware.Middleware) {
	s.middlewares.Add(middleware)
}
