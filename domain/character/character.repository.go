package character

type CharacterRepository interface {
	CharaCount() (int, error)
	FindChara(int) (string, error)
	CreateChara(Character) error
}
