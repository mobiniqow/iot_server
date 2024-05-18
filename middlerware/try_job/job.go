package try_job

import (
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
	Content          string
	Code             int16
	MessageTryNumber int
	TimeCounter      time.Duration
	IsSuccess        JobState
}
