package server

import (
	"fmt"
	"iot/device"
	"iot/message"
	"iot/middlerware"
	"iot/middlerware/try_job"
	"iot/server/handler"
	"net"

	"github.com/go-kit/kit/log"
)

type server struct {
	Port        int
	middlewares middlerware.Middlewares
	logger      log.Logger
}

func New(port int, logger log.Logger, middlerware middlerware.Middlewares) *server {
	return &server{
		Port:        port,
		middlewares: middlerware,
		logger:      logger,
	}
}

func (s *server) Run() {
	for _, middleware := range s.middlewares.Middleware {
		middleware.Controller()
	}

	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", s.Port))
	if err != nil {
		s.logger.Log("Error", err)
		return
	}

	if err != nil {
		s.logger.Log("Error:", err)
		return
	}
	defer listener.Close()
	defer s.logger.Log("server shutdown")
	// fmt.Println("Server is listening on port 8080")
	s.logger.Log("Server is listening on port 8080")

	deviceManager := device.GetInstanceManager(s.logger)

	validator := message.Validator{}

	decoder := message.Decoder{Logger: s.logger}

	for {
		// Accept incoming connections
		conn, err := listener.Accept()
		if err != nil {
			s.logger.Log("Error:", err)
			continue
		}
		newDevice := device.Device{Conn: conn, ClientID: conn.RemoteAddr().String()}
		deviceManager.Add(newDevice)
		handler := handler.Handler{Connection: conn, DeviceManager: deviceManager, Logger: s.logger, Validator: validator,
			Decoder: decoder, Device: newDevice, Middleware: &s.middlewares}
		handler.Start()
	}
}

func (s *server) Use(middleware try_job.TryJob) {
	s.middlewares.Add(&middleware)
}
