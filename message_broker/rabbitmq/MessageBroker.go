package rabbitmq

import (
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/wagslane/go-rabbitmq"
	"iot/device"
	"iot/message"
	"iot/utils"
)

type MessageBroker struct {
	Url           string
	Logger        log.Logger
	Producer      Producer
	Consumer      Consumer
	RabbitMQ      *rabbitmq.Conn
	deviceManager *device.Manager
}

func NewMessageBroker(url string, logger log.Logger) MessageBroker {
	deviceManager := device.GetInstanceManager(logger)
	producer := Producer{
		RoutingKey: "backend_routing_key",
		Exchange:   "backend_exchange",
		Logger:     logger,
		Queue:      "backend_queue",
	}

	consumer := Consumer{
		RoutingKey: "socket_server_routing_key",
		Exchange:   "socket_server_exchange",
		Queue:      "socket_server_queue",
		Handler: func(d rabbitmq.Delivery) rabbitmq.Action {
			logger.Log("consumed: %v", string(d.Body))
			sec := utils.StringToMap(string(d.Body))
			deviceId := sec["device_id"].(string)
			_type := sec["type"].(string)
			payload := sec["payload"].(string)
			date := sec["date"].(string)
			fmt.Printf("deviceId %s \n", deviceId)
			fmt.Printf("_type %s \n", _type)
			fmt.Printf("payload %s \n", payload)
			fmt.Printf("date %s \n", date)
			deviceManager.SendMessageWithDeviceId(deviceId, _type, payload, date)
			return rabbitmq.Ack
		},
		Logger: logger,
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
	c.Logger.Log("Send message to :", deviceId)
	c.Producer.SendMessage(deviceId, message)
}
