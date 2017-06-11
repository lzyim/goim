package main

import (
	"github.com/soloslee/goim/config"
	"github.com/soloslee/goim/server"
)

func start() {
	config := &config.Config{8000}
	server.Start(config)
}

var run = start

func main() {
	run()
}
