package server

import "sync"

var lock = &sync.Mutex{}

type server struct {
}

var singleInstance *server

// create singlethon server
func New() *server {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			singleInstance = &server{}

		}
	}
	return singleInstance
}
