package main

import (
	"fmt"
	"log"

	"github.com/Satish-Masa/CA-Tech-Dojo-Go/config"
	"github.com/Satish-Masa/CA-Tech-Dojo-Go/infrastructure"
	"github.com/Satish-Masa/CA-Tech-Dojo-Go/interfaces"
	"github.com/jinzhu/gorm"
)

func init() {
	config.Init()
}

func main() {
	tmp := "%s:%s@%s/%s"
	connect := fmt.Sprintf(tmp, config.Config.DbUser, config.Config.Password, config.Config.Tcp, config.Config.DbName)
	driver := config.Config.SQLDriver
	db, err := gorm.Open(driver, connect)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := infrastructure.NewUserRepository(db)
	rest := &interfaces.Rest{UserRepository: repo, GachaRepository: repo}
	_ = rest.Start()
}
