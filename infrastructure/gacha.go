package infrastructure

import (
	"github.com/Satish-Masa/CA-Tech-Dojo-Go/domain"
)

func CharaCount() (int, error) {
	db, err := ConnectDB()
	if err != nil {
		return -1, err
	}
	defer db.Close()
	var count int
	db.Table("character").Count(&count)

	return count, nil
}

func FindChara(num int) (string, error) {
	db, err := ConnectDB()
	if err != nil {
		return "", err
	}
	defer db.Close()
	var chara domain.Character
	db.Where("characterID=?", num).Find(&chara)

	return chara.Name, nil
}

func UpdateChar(chara domain.Character) error {
	db, err := ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close()
	err = db.Create(&chara).Error
	if err != nil {
		return err
	}
	return nil
}
