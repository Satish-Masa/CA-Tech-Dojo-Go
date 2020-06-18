package character

type Character struct {
	ID   int    `json: "id" gorm: "praimaly_key"`
	Name string `json: "name"`
}

func NewCharacter(name string) *Character {
	return &Character{
		Name: name,
	}
}
