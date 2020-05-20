package main

import (
	"github.com/Satish-Masa/CA-Tech-Dojo-Go/config"
	"github.com/Satish-Masa/CA-Tech-Dojo-Go/server"
)

func init() {
	config.Init()
}

func main() {
	server.Start()
}
