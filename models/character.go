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

	db.AutoMigrate(&UserCharacter{})
}
