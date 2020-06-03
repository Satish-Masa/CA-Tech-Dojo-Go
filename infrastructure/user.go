package infrastructure

import (
	"fmt"
	"log"

	application "github.com/Satish-Masa/CA-Tech-Dojo-Go/application/user"
	"github.com/Satish-Masa/CA-Tech-Dojo-Go/config"
	"github.com/Satish-Masa/CA-Tech-Dojo-Go/domain"
	"github.com/jinzhu/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	tmp := "%s:%s@%s/%s"
	connect := fmt.Sprintf(tmp, config.Config.DbUser, config.Config.Password, config.Config.Tcp, config.Config.DbName)
	driver := config.Config.SQLDriver
	db, err := gorm.Open(driver, connect)
	return db, err
}

func SaveUser(u *domain.User) error {
	db, err := ConnectDB()
	if err != nil {
		return err
	}
	db.Close()
	err := db.Create(&u).Error
	if err != nil {
		return err
	}

	return nil
}

func FindUser(token string) *application.UserGetResponce {
	db, err := ConnectDB()
	if err != nil {
		log.Println(err)
	}
	db.Close()
	resp := new(application.UserGetResponce)
	db.First(&resp, "name=?", token)
	return &resp
}

func UpdateUser(u *domain.User) error {
	db, err := ConnectDB()
	if err != nil {
		return err
	}
	db.Close()
	err := db.Model(&u).Where("token=?", u.Token).Update("name", u.Name).Error
	if err != nil {
		return err
	}
	return nil
}
