package server

import (
	"fmt"
	"iot/brodcaster"
	"iot/device"
	"iot/message"
	"iot/message_broker/rabbitmq"
	"iot/middlerware/try_job"
	"iot/server/handler"
	"net"

	"github.com/go-kit/kit/log"
)

type server struct {
	Port          int
	logger        log.Logger
	messageBroker rabbitmq.MessageBroker
	DeviceManager *device.Manager
	BroadCaster   *brodcaster.BroadCaster
}

func New(port int, logger log.Logger,
	messageBroker rabbitmq.MessageBroker, manager *device.Manager,
	broadCaster *brodcaster.BroadCaster,
) *server {
	return &server{
		Port:          port,
		logger:        logger,
		messageBroker: messageBroker,
		DeviceManager: manager,
		BroadCaster:   broadCaster,
	}
}

func (c *server) Run() {
	for _, middleware := range c.BroadCaster.MiddleWares.Middleware {
		middleware.Controller()
	}

	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", c.Port))
	if err != nil {
		c.logger.Log("Error", err)
		return
	}

	if err != nil {
		c.logger.Log("Error:", err)
		return
	}
	defer listener.Close()
	defer c.logger.Log("server shutdown")
	// fmt.Println("Server is listening on port 8080")
	c.logger.Log("Server is listening on port 8080")

	validator := message.Validator{}

	decoder := message.Decoder{Logger: c.logger}
	// rabbit mq consumer run
	go c.messageBroker.Run()
	for {
		// Accept incoming connections
		conn, err := listener.Accept()
		if err != nil {
			c.logger.Log("Error:", err)
			continue
		}
		newDevice := device.Device{Conn: conn, ClientID: conn.RemoteAddr().String()}

		c.DeviceManager.Add(newDevice)
		handler := handler.Handler{Connection: conn, DeviceManager: c.DeviceManager, Logger: c.logger,
			Validator: validator, Decoder: decoder, Device: newDevice, Middleware: c.BroadCaster.MiddleWares,
			//MessageBroker: rabbitmq.NewMessageBroker("amqp://guest:guest@localhost", c.logger)
			MessageBroker: c.messageBroker,
		}
		handler.Start()
	}
}

func (c *server) Use(middleware try_job.TryJob) {
	c.BroadCaster.MiddleWares.Add(&middleware)
}
