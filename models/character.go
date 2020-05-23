package models

type CharacterListResponce struct {
	Character UserCharacter
}

type UserCharacter struct {
	UserCharacterID string `json: "userCharacterID"`
	CharacterID     string `json: "characterID"`
	Name            string `json: "name"`
}
