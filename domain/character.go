package domain

type Character struct {
	UserCharacterID string `json: "userCharacterID"`
	CharacterID     string `json: "characterID"`
	Name            string `json: "name"`
}
