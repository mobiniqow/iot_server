package rabbitmq

import (
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/wagslane/go-rabbitmq"
	"iot/brodcaster"
	"iot/device"
	"iot/message"
	"iot/message_broker/gateway"
)

type MessageBroker struct {
	Url           string
	Logger        log.Logger
	Producer      Producer
	Consumer      Consumer
	RabbitMQ      *rabbitmq.Conn
	deviceManager *device.Manager
	Gateway       *gateway.Gateway
}

func NewMessageBroker(url string, logger log.Logger, _gateway *gateway.Gateway, deviceManager *device.Manager, brodCaster *brodcaster.BroadCaster) MessageBroker {
	producer := Producer{
		RoutingKey: "backend_routing_key",
		Exchange:   "backend_exchange",
		Logger:     logger,
		Queue:      "backend_queue",
	}
	_consumerHandler := ConsumerHandler{
		DeviceManager: deviceManager,
		Logger:        logger,
		Gateway:       _gateway,
		BroadCaster:   brodCaster,
	}
	consumer := Consumer{
		RoutingKey: "socket_server_routing_key",
		Exchange:   "socket_server_exchange",
		Queue:      "socket_server_queue",
		Handler:    _consumerHandler,
		Logger:     logger,
	}

	conn, _ := rabbitmq.NewConn(
		url,
		rabbitmq.WithConnectionOptionsLogging,
	)

	instance := MessageBroker{
		Url:           url,
		Logger:        logger,
		Producer:      producer,
		Consumer:      consumer,
		RabbitMQ:      conn,
		deviceManager: deviceManager,
	}
	return instance
}

func (c *MessageBroker) Run() {
	c.Logger.Log("Starting Consuming")
	c.Consumer.Run(c.RabbitMQ)
	defer c.Logger.Log("Ending Consuming...")
	defer c.RabbitMQ.Close()
}

func (c *MessageBroker) SendData(deviceId string, message message.Message) {
	fmt.Printf("Send message to : %v\r\n", deviceId)
	c.Producer.SendMessage(deviceId, message)
}
