package models

type CharacterListResponce struct {
	Character UserCharacter
}

type UserCharacter struct {
	UserCharacterID string `json: "userCharacterID"`
	CharacterID     string `json: "characterID"`
	Name            string `json: "name"`
}

func CreatUserCharacterTable() error {
	db, err := ConnectDB()
	if err != nil {
		return err
	}
	db.Close()
	if err = db.AutoMigrate(&UserCharacter{}).Error; err != nil {
		return err
	}

	db.Create(&UserCharacter{
		UserCharacterID: "",
		CharacterID:     "0001",
		Name:            "User01",
	})
	db.Create(&UserCharacter{
		UserCharacterID: "",
		CharacterID:     "0002",
		Name:            "User02",
	})
	db.Create(&UserCharacter{
		UserCharacterID: "",
		CharacterID:     "0003",
		Name:            "User03",
	})
	db.Create(&UserCharacter{
		UserCharacterID: "",
		CharacterID:     "0004",
		Name:            "User04",
	})
}
