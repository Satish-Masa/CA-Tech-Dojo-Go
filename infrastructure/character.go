package infrastructure

import (
	domainCharacter "github.com/Satish-Masa/CA-Tech-Dojo-Go/domain/character"
	Err "github.com/Satish-Masa/CA-Tech-Dojo-Go/error"
	"github.com/jinzhu/gorm"
)

type characterRepository struct {
	conn *gorm.DB
}

func NewCharacterRepository(conn *gorm.DB) domainCharacter.CharacterRepository {
	return &characterRepository{conn: conn}
}

func (c characterRepository) Count() (count int, err error) {
	err = c.conn.Table("characters").Count(&count).Error
	if err != nil {
		return -1, Err.ErrCount
	}
	return count, nil
}

func (c *characterRepository) Create(chara *domainCharacter.Character) error {
	err := c.conn.Create(&chara).Error
	if err != nil {
		return Err.ErrCreateChara
	}
	return nil
}
