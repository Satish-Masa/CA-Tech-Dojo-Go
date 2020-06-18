package infrastructure

import (
	"net/http"

	domainCharacter "github.com/Satish-Masa/CA-Tech-Dojo-Go/domain/character"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type characterRepository struct {
	conn *gorm.DB
}

func NewCharacterRepository(conn *gorm.DB) domainCharacter.CharacterRepository {
	return &characterRepository{conn: conn}
}

func (c characterRepository) Count() (count int, err error) {
	err = c.conn.Table("character").Count(&count).Error
	if err != nil {
		return -1, &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "failed to count characters",
		}
	}
	return count, nil
}

func (c characterRepository) Find(id int) (domainCharacter.Character, error) {
	var chara domainCharacter.Character
	err := c.conn.First(&chara, "id = ?", id).Error
	if err != nil {
		return domainCharacter.Character{}, &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "failed to find the character",
		}
	}
	return chara, nil
}

func (c characterRepository) Create(chara *domainCharacter.Character) error {
	err := c.conn.Create(&chara).Error
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "failed to save the chara",
		}
	}
	return nil
}
