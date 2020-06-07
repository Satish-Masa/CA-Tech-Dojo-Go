package repository

import (
	"github.com/Satish-Masa/CA-Tech-Dojo-Go/application/user"
	"github.com/Satish-Masa/CA-Tech-Dojo-Go/domain"
)

type UserRepository interface {
	SaveUser(domain.User) error
	FindUser(domain.User) user.UserGetResponce
	UpdateUser(name, token string) error
}
