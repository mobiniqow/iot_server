package gateway

import (
	"errors"
	"fmt"
	"github.com/go-kit/kit/log"
	"iot/message"
	"iot/strategy"
	"iot/utils"
)

type Gateway struct {
	strategy map[string]strategy.Strategy
	logger   log.Logger
}

func NewGateway(logger log.Logger) *Gateway {
	return &Gateway{
		strategy: make(map[string]strategy.Strategy),
		logger:   logger,
	}
}

func (c *Gateway) MessageBroker(data string) (message.Message, error) {
	c.logger.Log("Gateway", "Input", "deviceId")

	strategyCode := c.GetStrategyCode(data)
	if strategyCode[0] == 'W' {
		strategyCode = fmt.Sprintf("R%c", strategyCode[1])
	}
	println("strategyCode: ", strategyCode)
	val, ok := c.strategy[strategyCode]
	if ok {
		return val.MessageBroker(data)
	}
	return message.Message{}, errors.New("strategy not found")
}

func (c *Gateway) ClientHandler(data string) (message.Message, error) {
	c.logger.Log("Gateway", "Input", "deviceId")
	if len(data) >= 2 {
		strategyCode := data[:2]
		if strategyCode[0] == 'W' {
			strategyCode = fmt.Sprintf("R%c", strategyCode[1])
			println("strategyCode2: ", fmt.Sprintf("R%c", strategyCode[1]))
		}
		println("strategyCode: ", strategyCode)
		val, ok := c.strategy[strategyCode]
		if ok {
			return val.ClientHandler(data)
		}
		return message.Message{}, errors.New("strategy not found")
	}
	return message.Message{}, errors.New("data a bit is wrong")
}

func (c *Gateway) AddStrategy(strategy strategy.Strategy) {
	c.logger.Log("Gateway", "AddStrategy", "strategy", strategy)
	c.strategy[strategy.GetCode()] = strategy
}

func (c *Gateway) GetStrategyCode(data string) string {
	if data[0] == 'W' {
		data = fmt.Sprintf("R%S", data[1])
	}
	dataString := data
	dataMap := utils.StringToMap(dataString)
	_type := dataMap["type"].(string)
	return _type
}
