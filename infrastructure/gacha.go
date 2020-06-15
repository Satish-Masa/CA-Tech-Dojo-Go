package infrastructure

import (
	"net/http"

	domainCharacter "github.com/Satish-Masa/CA-Tech-Dojo-Go/domain/character"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type gachaRepository struct {
	conn *gorm.DB
}

func NewGachaRepository(conn *gorm.DB) domainCharacter.CharacterRepository {
	return &gachaRepository{conn: conn}
}

func (g *gachaRepository) CharaCount() (count int, err error) {
	err = g.conn.Table("character").Count(&count).Error
	if err != nil {
		return -1, &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "failed to count characters",
		}
	}
	return count, nil
}

func (g *gachaRepository) FindChara(num int) (string, error) {
	var chara domainCharacter.Character
	err := g.conn.Where("characterID=?", num).Find(&chara).Error
	if err != nil {
		return "", &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "failed to find the character",
		}
	}
	return chara.Name, nil
}

func (g *gachaRepository) CreateChara(chara domainCharacter.Character) error {
	err := g.conn.Create(&chara).Error
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "failed to create the character",
		}
	}
	return nil
}
