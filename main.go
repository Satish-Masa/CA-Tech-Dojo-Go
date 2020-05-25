package main

import (
	"log"

	"github.com/Satish-Masa/CA-Tech-Dojo-Go/config"
	"github.com/Satish-Masa/CA-Tech-Dojo-Go/server"
)

func init() {
	config.Init()
	err := models.CreateUserTable() if err != nil {
		log.Println(err)
	}
	err = models.CreatUserCharacterTable() if err != nil {
		log.Println(err)
	}
}

func main() {
	server.Start()
}
