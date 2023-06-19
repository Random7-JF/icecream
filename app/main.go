package main

import (
	"github.com/Random-7/icecream/app/config"
	"github.com/Random-7/icecream/app/server"
)

func main() {

	var appConfig config.App
	appConfig.IP = "127.0.0.0"
	appConfig.Port = "4000"

	server.Serve(&appConfig)
}
