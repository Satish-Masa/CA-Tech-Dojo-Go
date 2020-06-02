package application

import (
	"log"

	"github.com/Satish-Masa/CA-Tech-Dojo-Go/infrastructure"
	"github.com/dgrijalva/jwt-go"
)

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

func SaveUser(name, token string) error {
	return infrastructure.SaveUser(name, token)
}

func SearchUser(token string) *UserGetResponce {
	return infrastructure.SearchUser(token)
}

func UpdateUser(name, token string) error {
	return infrastructure.UpdateUser(name, token)
}
