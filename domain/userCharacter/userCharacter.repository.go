package userCharacter

type UserCharacterRepository interface {
	Create(UserCharacter) error
	Find(int) (string, error)
	FindAll(int) ([]UserCharacter, error)
}
