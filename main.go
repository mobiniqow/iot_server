package main

import (
	"iot/server"
	"iot/server/middleware"
)

type MyEnum string

const (
	Foo MyEnum = "VV"
	Bar MyEnum = "BB"
)

func main() {
	PORT := 8080
	server := server.New(PORT)
	server.Use(&middleware.TryJob{})
	server.Run()

}
