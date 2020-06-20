package infrastructure

import (
	domainUser "github.com/Satish-Masa/CA-Tech-Dojo-Go/domain/user"
	Err "github.com/Satish-Masa/CA-Tech-Dojo-Go/error"
	"github.com/jinzhu/gorm"
)

type userRepository struct {
	conn *gorm.DB
}

type UserCharacter struct {
	UserCharacterID string `json: "userCharacterID"`
	CharacterID     int    `json: "characterID"`
	Name            string `json: "name"`
}

type CharacterListResponse struct {
	Characters []UserCharacter `json: "characters"`
}

func NewUserRepository(conn *gorm.DB) domainUser.UserRepository {
	return &userRepository{conn: conn}
}

func (i *userRepository) Save(u *domainUser.User) error {
	err := i.conn.Create(&u).Error
	if err != nil {
		return Err.ErrCreateUser
	}

	return nil
}

func (i *userRepository) Find(id int) (domainUser.User, error) {
	var user domainUser.User
	err := i.conn.First(&user, "id = ?", id).Error
	if err != nil {
		return domainUser.User{}, Err.ErrFindUser
	}
	return user, nil
}

func (i *userRepository) Update(name string, id int) error {
	err := i.conn.Model(&domainUser.User{}).Where("id=?", id).Update("name", name).Error
	if err != nil {
		return Err.ErrUpdateUser
	}
	return nil
}
