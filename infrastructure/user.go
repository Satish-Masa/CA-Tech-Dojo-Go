package infrastructure

import (
	"net/http"

	domainUser "github.com/Satish-Masa/CA-Tech-Dojo-Go/domain/user"
	"github.com/Satish-Masa/CA-Tech-Dojo-Go/interfaces"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type userRepository struct {
	conn *gorm.DB
}

type UserGetResponce struct {
	Name string `json: "name"`
}

type UserCreatResponse struct {
	Token string `json: "token"`
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
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "failed to save the user",
		}
	}

	return nil
}

func (i *userRepository) Find(id int) (*UserGetResponce, error) {
	resp := new(UserGetResponce)
	err := i.conn.First(&resp, "name=?", id).Error
	if err != nil {
		return &UserGetResponce{}, &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "failed to find the user",
		}
	}
	return resp, nil
}

func (i *userRepository) Update(u *interfaces.UserUpdateRequest, id int) error {
	err := i.conn.Model(&domainUser.User{}).Where("id=?", id).Update("name", u.Name).Error
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "failed to update the user",
		}
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
