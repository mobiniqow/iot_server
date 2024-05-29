package rabbitmq

import (
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/wagslane/go-rabbitmq"
	"iot/message"
)

type Producer struct {
	Queue      string
	Exchange   string
	RoutingKey string
	Logger     log.Logger
}

func (c *Producer) SendMessage(deviceID string, message message.Message) error {

	conn, err := rabbitmq.NewConn(
		"amqp://guest:guest@localhost",
		rabbitmq.WithConnectionOptionsLogging,
	)

	if err != nil {
		c.Logger.Log(err)
	}

	defer conn.Close()

	publisher, err := rabbitmq.NewPublisher(
		conn,
		rabbitmq.WithPublisherOptionsLogging,
		rabbitmq.WithPublisherOptionsExchangeName(c.Exchange),
		rabbitmq.WithPublisherOptionsExchangeDeclare,
	)

	if err != nil {
		c.Logger.Log(err)
	}

	defer publisher.Close()

	fmt.Printf("c.RoutingKey %v \n", c.RoutingKey)

	err = publisher.Publish(
		// inja status ro bapayload ezafekarda var dar ebtedash device id ro dadam => devoceId:type+payload
		append([]byte(deviceID+":"), append(message.Type, message.Payload[:]...)[:]...),
		[]string{c.RoutingKey},
		rabbitmq.WithPublishOptionsContentType("application/json"),
		rabbitmq.WithPublishOptionsExchange(c.Exchange),
	)

	if err != nil {
		c.Logger.Log(err)
	}

	return nil

}
