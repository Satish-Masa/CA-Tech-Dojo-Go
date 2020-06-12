package character

type CharacterRepository interface {
	CharaCount() (int, error)
	FindChara(int) (string, error)
	UpdateChara(Character) error
}
