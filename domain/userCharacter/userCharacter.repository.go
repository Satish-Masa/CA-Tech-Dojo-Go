package userCharacter

type UserCharacterRepository interface {
	Create(UserCharacter) error
}
