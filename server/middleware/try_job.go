package middleware

import "net"

type TryJob struct {
}

func (tj *TryJob) Output(con net.Conn, err error) (net.Conn, error) {
	if err != nil {
		return con, err
	}
	print("con is %v", con)
	return con, nil
}

func (tj *TryJob) Input(con net.Conn, err error) (net.Conn, error) {
	return con, err
}
