package try_job

import (
	"fmt"
	"iot/device"
	"iot/message"
	"iot/utils"
	"net"
	"time"
)

// period time between users period job
const PERIOD = 100 * time.Millisecond

var JOB_QUEUE int16 = 0

const Name = "TRY_JOB"
const CODE = 0xFAFA

type TryJob struct {
	TryNumber int8
	SleepTime time.Duration
	Jobs      map[string]Job
}

// todo bayad ersal konam va sare har ersal ye shomare bezanam va vaghty packet ersal shode hast va dobare ersal kard
// todo betone packet ro dobare ersal kona az counter 0 beshe
// todo check konam agevice to device manager nist job ro pak konam

func (c *TryJob) Controller() {
	go func() {
		for {
			for k, v := range c.Jobs {
				if v.State == SUSPENDED {
					dm := device.GetInstanceManagerWithoutLogger()
					job := c.Jobs[k]
					job.MessageTryNumber = job.MessageTryNumber + 1
					c.Jobs[k] = job
					dm.SendMessage(v.Conn, &job.Data)
					if job.MessageTryNumber >= c.TryNumber {
						job.State = END
						c.Jobs[k] = job
						// inja shayad end haro pak konam
					}
				}
			}
			time.Sleep(c.SleepTime)
			//print("Sleep time: ", c.SleepTime, "Seconds", "\n")
		}
	}()
}

func (c *TryJob) Output(con *net.Conn, data *message.Message) error {
	key := utils.JobKeyGenerator(*con, *data)
	_, ok := c.Jobs[key]
	if ok {
		messageCode := fmt.Sprintf("%04X%02X", c.Jobs[key].Code, c.Jobs[key].MessageTryNumber)
		extention := message.Extention{Name: Name, Code: messageCode}
		data.Extentions = append(data.Extentions, extention)
	} else {
		JOB_QUEUE++
		job := Job{
			Conn:             *con,
			Data:             *data,
			Code:             JOB_QUEUE,
			MessageTryNumber: 0,
			TimeCounter:      c.SleepTime,
			State:            SUSPENDED,
		}
		c.Jobs[key] = job
		messageCode := fmt.Sprintf("%04X%02X", job.Code, job.MessageTryNumber)
		extention := message.Extention{Name: Name, Code: messageCode}
		data.Extentions = append(data.Extentions, extention)
	}
	return nil
}

func (c *TryJob) Input(con *net.Conn, message *message.Message) error {
	return nil
}
