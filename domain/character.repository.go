package domain

import (
	"github.com/Satish-Masa/CA-Tech-Dojo-Go/domain"
)

type CharacterRepository interface {
	CharaCount() (int, error)
	FindChara(int) (string, error)
	UpdateChara(domain.Character) error
}
