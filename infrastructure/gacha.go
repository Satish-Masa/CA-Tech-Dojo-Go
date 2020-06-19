package infrastructure

import (
	"net/http"

	domainCharacter "github.com/Satish-Masa/CA-Tech-Dojo-Go/domain/character"
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

func (c *gachaRepository) Find(id int) (string, error) {
	var chara domainCharacter.Character
	err := c.conn.First(&chara, "id = ?", id).Error
	if err != nil {
		return "", &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "failed to find characters",
		}
	}
	return chara.Name, nil
}

func (c *gachaRepository) FindAll(id int) ([]domainUserCharacter.UserCharacter, error) {
	var charas []domainUserCharacter.UserCharacter
	err := c.conn.Where("user_character_id = ?", id).Find(&charas).Error
	if err != nil {
		return charas, err
	}

	return charas, nil
}
