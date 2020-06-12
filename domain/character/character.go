package character

type Character struct {
	UserCharacterID string `json: "userCharacterID"`
	CharacterID     int    `json: "characterID"`
	Name            string `json: "name"`
}

func NewCharacter(uid string, cid int, name string) *Character {
	return &Character{
		UserCharacterID: uid,
		CharacterID:     cid,
		Name:            name,
	}
}
