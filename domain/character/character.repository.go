package character

type CharacterRepository interface {
	Count() (int, error)
	Create(*Character) error
}
