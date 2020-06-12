package infrastructure

import (
	domainCharacter "github.com/Satish-Masa/CA-Tech-Dojo-Go/domain/character"
	"github.com/jinzhu/gorm"
)

type gachaRepository struct {
	conn *gorm.DB
}

func (g *gachaRepository) CharaCount() (count int, err error) {
	g.conn.Table("character").Count(&count)
	return count, nil
}

func (g *gachaRepository) FindChara(num int) (string, error) {
	var chara domain.Character
	g.conn.Where("characterID=?", num).Find(&chara)
	return chara.Name, nil
}

func (g *gachaRepository) UpdateChara(chara domainCharacter.Character) error {
	err := g.conn.Create(&chara).Error
	if err != nil {
		return err
	}
	return nil
}
