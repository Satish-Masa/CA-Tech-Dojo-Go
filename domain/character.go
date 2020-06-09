package domain

type Character struct {
	UserCharacterID string `json: "userCharacterID"`
	CharacterID     int    `json: "characterID"`
	Name            string `json: "name"`
}
