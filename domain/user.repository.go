package domain

import (
	"github.com/Satish-Masa/CA-Tech-Dojo-Go/domain"
	"github.com/Satish-Masa/CA-Tech-Dojo-Go/infrastructure"
	"github.com/Satish-Masa/CA-Tech-Dojo-Go/interfaces"
)

type UserRepository interface {
	Save(*domain.User) (infrastructure.UserCreatResponse, error)
	Find(*domain.User) infrastructure.UserGetResponce
	Update(*interfaces.UserUpdateRequest) error
}
