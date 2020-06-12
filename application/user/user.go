package user

import (
	"github.com/Satish-Masa/CA-Tech-Dojo-Go/domain"
	"github.com/Satish-Masa/CA-Tech-Dojo-Go/domain/repository"
	"github.com/Satish-Masa/CA-Tech-Dojo-Go/infrastructure"
	"github.com/Satish-Masa/CA-Tech-Dojo-Go/interfaces"
)

type UserApplication struct {
	Repository repository.UserRepository
}

type UserCharacter struct {
	UserCharacterID string `json: "userCharacterID"`
	CharacterID     int    `json: "characterID"`
	Name            string `json: "name"`
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
