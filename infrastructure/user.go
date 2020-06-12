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

func (i *userRepository) Save(u *domain.User) (*UserCreatResponse, error) {
	err := i.conn.Create(&u).Error
	if err != nil {
		return nil, err
	}

	return &UserCreatResponse{Token: u.Token}, nil
}

func (i *userRepository) Find(u *domain.User) *UserGetResponce {
	resp := new(UserGetResponce)
	i.conn.First(&resp, "name=?", u.Token)
	return resp
}

func (i *userRepository) Update(u *interfaces.UserUpdateRequest) error {
	err := i.conn.Model(&u).Where("token=?", u.Token).Update("name", u.Name).Error
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
