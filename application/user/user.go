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

func (a UserApplication) SaveUser(u *domain.User) error {
	return a.Repository.Save(u)
}

func (a UserApplication) FindUser(uid int) infrastructure.UserGetResponce {
	return a.Repository.Find(uid)
}

func (a UserApplication) UpdateUser(u *interfaces.UserUpdateRequest, id int) error {
	return a.Repository.Update(u, id)
}

/* func (a UserApplication) GetList(u domain.User) (CharacterListResponse, error) {
	return infrastructure.FindChara(u.Token)
} */
