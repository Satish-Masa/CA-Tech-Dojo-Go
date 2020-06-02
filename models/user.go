package models

import (
	"fmt"
	"log"

	"github.com/Satish-Masa/CA-Tech-Dojo-Go/config"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type User struct {
	Name  string `json: "name"`
	Token string `json: "token"`
}

type UserCreatRequest struct {
	Name string `json: "name"`
}

type UserCreatResponse struct {
	Token string `json: "token"`
}

type UserGetResponce struct {
	Name string `json: "name"`
}

type UserUpdateRequest struct {
	Name  string `json: "name"`
	Token string `json: "token"`
}

func ConnectDB() (*gorm.DB, error) {
	tmp := "%s:%s@%s/%s"
	connect := fmt.Sprintf(tmp, config.Config.DbUser, config.Config.Password, config.Config.Tcp, config.Config.DbName)
	driver := config.Config.SQLDriver
	db, err := gorm.Open(driver, connect)
	return db, err
}

func CreateUserTable() error {
	db, err := ConnectDB()
	if err != nil {
		return err
	}
	db.Close()
	if err := db.AutoMigrate(&User{}).Error; err != nil {
		return err
	}
	return nil
}

func creatToken(name string) (string, error) {
	var err error
	secret := "secret"
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": name,
		"iss":  "__init__",
	})
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Fatal(err)
	}
	return tokenString, nil
}

func FetchToken(name string) *UserCreatResponse {
	token, err := creatToken(name)
	if err != nil {
		log.Println(err)
	}
	resp := new(UserCreatResponse)
	resp.Token = token

	return &resp
}

func NewUser(name, token string) *User {
	return &User{
		name,
		token,
	}
}

func SaveUser(name, token string) error {
	u := NewUser(name, token)
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

func SearchUser(token string) *UserGetResponce {
	db, err := ConnectDB()
	if err != nil {
		log.Println(err)
	}
	db.Close()
	resp := new(UserGetResponce)
	db.First(&resp, "name=?", token)
	return &resp
}

func UpdateUser(name, token string) error {
	db, err := ConnectDB()
	if err != nil {
		return err
	}
	db.Close()
	var u User
	err := db.Model(&u).Where("token=?", token).Update("name", name).Error
	if err != nil {
		return err
	}
	return nil
}
