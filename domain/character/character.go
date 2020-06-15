package character

type Character struct {
	UserCharacterID int    `json: "userCharacterID"`
	CharacterID     int    `json: "characterID"`
	Name            string `json: "name"`
}

func NewCharacter(uid int, cid int, name string) *Character {
	return &Character{
		UserCharacterID: uid,
		CharacterID:     cid,
		Name:            name,
	}
}
