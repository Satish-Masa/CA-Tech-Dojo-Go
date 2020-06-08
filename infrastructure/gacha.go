package infrastructure

import (
	"github.com/Satish-Masa/CA-Tech-Dojo-Go/application/gacha"
	"github.com/Satish-Masa/CA-Tech-Dojo-Go/domain"
)

type gachaRepository struct {
	ConnectDB() (*gorm.DB, error)
}

func ConnectDB() (*gorm.DB, error) {
	tmp := "%s:%s@%s/%s"
	connect := fmt.Sprintf(tmp, config.Config.DbUser, config.Config.Password, config.Config.Tcp, config.Config.DbName)
	driver := config.Config.SQLDriver
	db, err := gorm.Open(driver, connect)
	return db, err
}

func (g *gachaRepository) CharaCount() (int, error) {
	db, err := g.ConnectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	var count int
	db.Table("character").Count(&count)

	return count, nil
}

func (g *gachaRepository) FindChara(n int) (string, error) {
	db, err := g.ConnectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	var result gacha.GachaResult
	db.First(&result, "name=?", characterID)

	return result.Name, nil
}

func (g *gachaRepository) UpdateChar(chara domain.Character) error {
	db, err := g.ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close()
	err := db.Create(&chara).Error
	if err != nil {
		return err
	}
	return nil
}
