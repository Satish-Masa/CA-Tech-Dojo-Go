package character

import domainCharacter "github.com/Satish-Masa/CA-Tech-Dojo-Go/domain/character"

type CharacterApplication struct {
	Repository domainCharacter.CharacterRepository
}

type CharaCreateRequest struct {
	Name string `json: "name"`
}

func (r CharacterApplication) CreateChara(chara *domainCharacter.Character) error {
	return r.Repository.Create(chara)
}

func (r CharacterApplication) FindChara(id int) (domainCharacter.Character, error) {
	return r.Repository.Find(id)
}

func (r CharacterApplication) CountChara() (int, error) {
	return r.Repository.Count()
}
