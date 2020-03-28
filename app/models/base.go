package models

import (
	"CA-Tech-Dojo-Go/config"
	"fmt"

	"github.com/jinzhu/gorm"
)

func ConnectDB() *gorm.DB {
	tmp := "%s:%s@%s/%s"
	connect := fmt.Sprintf(tmp, config.Config.DbUser, config.Config.Password, config.Config.Tcp, config.Config.DbName)
	driver := config.Config.SQLDriver
	db, _ := gorm.Open(driver, connect)
	autoMigration(db)
	return db
}

func autoMigration(db *gorm.DB) {
	db.AutoMigrate(&User{})
}
