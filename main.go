package main

import (
	"fmt"
	"log"

	"github.com/Satish-Masa/CA-Tech-Dojo-Go/config"
	domainChara "github.com/Satish-Masa/CA-Tech-Dojo-Go/domain/character"
	domainUser "github.com/Satish-Masa/CA-Tech-Dojo-Go/domain/user"
	"github.com/Satish-Masa/CA-Tech-Dojo-Go/infrastructure"
	"github.com/Satish-Masa/CA-Tech-Dojo-Go/interfaces"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	config.Init()
}

func main() {
	tmp := "%s:%s@/%s?charset=utf8&parseTime=True&loc=Local"
	connect := fmt.Sprintf(tmp, config.Config.DbUser, config.Config.Password, config.Config.DbName)
	driver := config.Config.SQLDriver
	db, err := gorm.Open(driver, connect)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.AutoMigrate(&domainUser.User{})
	db.AutoMigrate(&domainChara.Character{})

	user := infrastructure.NewUserRepository(db)
	gacha := infrastructure.NewGachaRepository(db)
	rest := &interfaces.Rest{UserRepository: user, GachaRepository: gacha}
	rest.Start()
}
