package infrastructure

import "github.com/Satish-Masa/CA-Tech-Dojo-Go/application/gacha"

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
