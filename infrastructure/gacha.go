package infrastructure

import (
	"github.com/Satish-Masa/CA-Tech-Dojo-Go/application/gacha"
	"github.com/Satish-Masa/CA-Tech-Dojo-Go/domain"
)

func CharaCount() int {
	db, err := ConnectDB()
	if err != nil {
		return err
	}
	db.Close()
	count := db.Table("character").Count()

	return count
}

func FindChara(int n) string {
	db, err := ConnectDB()
	if err != nil {
		return err
	}
	db.Close()
	var result gacha.GachaResult
	db.First(&result, "name=?", characterID)

	return result.Name
}

func UpdateChar(chara domain.Character) error {
	db, err := ConnectDB()
	if err != nil {
		return err
	}
	db.Close()
	err := db.Create(&chara).Error
	if err != nil {
		return err
	}
	return nil
}
