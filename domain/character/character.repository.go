package character

type CharacterRepository interface {
	Count() (int, error)
	Find(int) (Character, error)
	Create(*Character) error
}
