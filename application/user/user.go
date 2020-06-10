package user

import (
	"log"

	"github.com/Satish-Masa/CA-Tech-Dojo-Go/domain"
	"github.com/Satish-Masa/CA-Tech-Dojo-Go/domain/repository"
	"github.com/Satish-Masa/CA-Tech-Dojo-Go/infrastructure"
	"github.com/Satish-Masa/CA-Tech-Dojo-Go/interfaces"
	"github.com/dgrijalva/jwt-go"
)

type UserApplication struct {
	Repository repository.UserRepository
}

type UserCharacter struct {
	UserCharacterID string `json: "userCharacterID"`
	CharacterID     int    `json: "characterID"`
	Name            string `json: "name"`
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

func FetchToken(u *interfaces.UserCreatRequest) (token string, err error) {
	token, err = creatToken(u.Name)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (a UserApplication) SaveUser(u *domain.User) (interfaces.UserCreatRequest, error) {
	return a.Repository.Save(u)
}

func (a UserApplication) FindUser(u *domain.User) infrastructure.UserGetResponce {
	return a.Repository.Find(u)
}

func (a UserApplication) UpdateUser(u *interfaces.UserUpdateRequest) error {
	return a.Repository.Update(u)
}

/* func (a UserApplication) GetList(u domain.User) (CharacterListResponse, error) {
	return infrastructure.FindChara(u.Token)
} */
