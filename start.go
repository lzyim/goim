package main

import (
	"github.com/soloslee/goim/config"
	"github.com/soloslee/goim/server"
)

func main() {
	config := &config.Config{8000}
	server.Start(config)
}
