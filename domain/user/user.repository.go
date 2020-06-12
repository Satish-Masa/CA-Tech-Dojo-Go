package user

import (
	"github.com/Satish-Masa/CA-Tech-Dojo-Go/infrastructure"
	"github.com/Satish-Masa/CA-Tech-Dojo-Go/interfaces"
)

type UserRepository interface {
	Save(*User) (infrastructure.UserCreatResponse, error)
	Find(*User) infrastructure.UserGetResponce
	Update(*interfaces.UserUpdateRequest) error
}
