package infrastructure

import (
	"net/http"

	domainUserCharacter "github.com/Satish-Masa/CA-Tech-Dojo-Go/domain/userCharacter"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type gachaRepository struct {
	conn *gorm.DB
}

func NewGachaRepository(conn *gorm.DB) domainUserCharacter.UserCharacterRepository {
	return &gachaRepository{conn: conn}
}

func (g *gachaRepository) Create(chara domainUserCharacter.UserCharacter) error {
	err := g.conn.Create(&chara).Error
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "failed to create the userCharacter",
		}
	}
	return nil
}
