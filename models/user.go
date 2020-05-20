package models

import (
	"log"

	"github.com/dgrijalva/jwt-go"
)

type User struct {
	Name  string
	Token string
}

type UserCreatRequest struct {
	Name string
}

type UserCreatResponse struct {
	Token string
}

type UserGetResponce struct {
	Name string
}

type UserUpdateRequest struct {
	Name string
}

func CreatToken(name string) (string, error) {
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

func (u *User) NewUser(name string, token string) *User {
	u.Name = name
	u.Token = token
	return &u
}
