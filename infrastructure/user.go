package infrastructure

import (
	"fmt"
	"log"

	"github.com/Satish-Masa/CA-Tech-Dojo-Go/application/user"
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

func Save(u *domain.User) error {
	db, err := ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close()
	err = db.Create(&u).Error
	if err != nil {
		return err
	}

	return nil
}

func Find(token string) *user.UserGetResponce {
	db, err := ConnectDB()
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	resp := new(user.UserGetResponce)
	db.First(&resp, "name=?", token)
	return resp
}

func Update(u *domain.User) error {
	db, err := ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close()
	err = db.Model(&u).Where("token=?", u.Token).Update("name", u.Name).Error
	if err != nil {
		return err
	}
	return nil
}

/* func FindChara(token string) user.CharacterListResponse {
	db, err := ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	var count int
	db.Model(domain.Character).Where("token=?", token).Count(&count)

	for i := 0; i < count; i++ {

	}
} */
