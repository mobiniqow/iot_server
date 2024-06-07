package gateway

import (
	"errors"
	"github.com/go-kit/kit/log"
	"iot/message"
	"iot/message_broker/strategy"
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

func (c *Gateway) Input(data []byte) (message.Message, error) {
	c.logger.Log("Gateway", "Input", "deviceId")
	strategyCode := c.GetStrategyCode(data)
	println("strategyCode: ", strategyCode)
	val, ok := c.strategy[strategyCode]
	if ok {
		return val.Input(data)
	}
	return message.Message{}, errors.New("strategy not found")
}

func (c *Gateway) Output(message message.Message, deviceId string) {
	c.logger.Log("Gateway", "Output", message.Type, "deviceId", deviceId)
}

func (c *Gateway) AddStrategy(strategy strategy.Strategy) {
	c.logger.Log("Gateway", "AddStrategy", "strategy", strategy)
	c.strategy[strategy.GetCode()] = strategy
}
func (c *Gateway) GetStrategyCode(data []byte) string {
	dataString := string(data)
	dataMap := utils.StringToMap(dataString)
	_type := dataMap["type"].(string)
	return _type
}
