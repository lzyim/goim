package main

import (
	"goim/config"
	"goim/server"
)

func main() {
	config := &config.Config{8000}
	server.Start(config)
}
