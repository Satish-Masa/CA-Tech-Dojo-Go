package infrastructure

import (
	"github.com/Satish-Masa/CA-Tech-Dojo-Go/application/user"
	domainUser "github.com/Satish-Masa/CA-Tech-Dojo-Go/domain/user"
	"github.com/Satish-Masa/CA-Tech-Dojo-Go/interfaces"
	"github.com/jinzhu/gorm"
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

type CharacterListResponse struct {
	Characters []user.UserCharacter `json: "characters"`
}

func NewUserRepository(conn *gorm.DB) domainUser.UserRepository {
	return &userRepository{conn: conn}
}

func (i *userRepository) Save(u *domainUser.User) error {
	err := i.conn.Create(&u).Error
	if err != nil {
		return err
	}

	return nil
}

func (i *userRepository) Find(id int) *UserGetResponce {
	resp := new(UserGetResponce)
	i.conn.First(&resp, "name=?", id)
	return resp
}

func (i *userRepository) Update(u *interfaces.UserUpdateRequest, id int) error {
	err := i.conn.Model(&domainUser.User).Where("id=?", id).Update("name", u.Name).Error
	if err != nil {
		return err
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
