package userCharacter

type UserCharacter struct {
	UserCharacterID int    `json: "userCharacterID"`
	CharacterID     int    `json: "characterID"`
	Name            string `json: "name"`
}

func NewCharacter(uid int, cid int, name string) *UserCharacter {
	return &UserCharacter{
		UserCharacterID: uid,
		CharacterID:     cid,
		Name:            name,
	}
}
