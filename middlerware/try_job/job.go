package try_job

import (
	"iot/message"
	"net"
	"time"
)

type JobState int

const (
	SUSPENDED JobState = iota
	SUCCESS
	FAILED
	END
)

type Job struct {
	Conn             net.Conn
	Data             message.Message
	Code             int16
	MessageTryNumber int8
	TimeCounter      time.Duration
	State            JobState
}
