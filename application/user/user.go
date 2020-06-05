package user

import (
	"log"

	"github.com/Satish-Masa/CA-Tech-Dojo-Go/domain"
	"github.com/Satish-Masa/CA-Tech-Dojo-Go/infrastructure"
	"github.com/dgrijalva/jwt-go"
)

type UserApplication struct {
	infra infrastructure.UserInfrastracture
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

func FetchToken(u *UserCreatRequest) (*UserCreatResponse, error) {
	token, err := creatToken(u.Name)
	if err != nil {
		return nil, err
	}

	return &UserCreatResponse{
		Token: token,
	}, nil
}

func (a *UserApplication) SaveUser(u *domain.User) error {
	return a.infra.Save(u)
}

func (a *UserApplication) FindUser(u *domain.User) *UserGetResponce {
	return a.infra.Find(u.Token)
}

func (a *UserApplication) UpdateUser(name, token string) error {
	u, err := domain.NewUser(name, token)
	if err != nil {
		return err
	}
	return a.infra.Update(u)
}
