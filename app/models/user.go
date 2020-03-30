package models

import _ "github.com/go-sql-driver/mysql"

type User struct {
	Name  string `json:"name"`
	Token string `gorm: "primary_key"`
}

type UserCreatReqest struct {
	Name string
}

type UserCreatResponse struct {
	Token string
}

type UserGetResponse struct {
	Name string
}

type UserUpdateRequest struct {
	Name string
}

func CreatUser(name string, token string) {
	sql := ConnectDB()
	defer sql.Close()
	var u User
	u.Name = name
	u.Token = token
	sql.Create(&u)
}

func GetUser(token string) string {
	sql := ConnectDB()
	defer sql.Close()
	var u User
	name := sql.First(&u, "token = ?", token).Dialect().GetName()
	return name
}
