package try_job

import (
	"net"
	"time"
)

// period time between users period job
const PERIOD = 100 * time.Millisecond

var JOB_QUEUE int16 = 0

type TryJob struct {
	TryNumber int8
	SleepTime time.Duration
	Jobs      map[string]Job
}

// todo bayad ersal konam va sare har ersal ye shomare bezanam va vaghty packet ersal shode hast va dobare ersal kard betone packet ro dobare ersal kona az counter 0 beshe

func (c *TryJob) Controller() {
	go func() {
		for {
			for k, v := range c.Jobs {
				println(k)
				println(333)
				println(v)
			}
			time.Sleep(PERIOD)
		}
	}()
}

func (c *TryJob) Output(con net.Conn, err error, data string) (net.Conn, error) {
	if err != nil {
		return con, err
	}

	key := JobKeyGenerator(con, data)
	_, ok := c.Jobs[key]
	if ok {
		print("job already exists")
	} else {
		print("job does not exist")
		JOB_QUEUE++
		job := Job{
			Conn:             con,
			Content:          data,
			Code:             JOB_QUEUE,
			MessageTryNumber: 0,
			TimeCounter:      c.SleepTime,
			IsSuccess:        SUSPENDED,
		}
		c.Jobs[key] = job
	}
	return con, nil
}

func (c *TryJob) Input(con net.Conn, err error, data string) (net.Conn, error) {
	if err != nil {
		return con, err
	}
	return con, err
}
