package gacha

import (
	"math/rand"
	"net/http"
	"time"

	domainCharacter "github.com/Satish-Masa/CA-Tech-Dojo-Go/domain/character"
	domainUserCharacter "github.com/Satish-Masa/CA-Tech-Dojo-Go/domain/userCharacter"
	"github.com/labstack/echo/v4"
)

type GachaApplication struct {
	Repository      domainUserCharacter.UserCharacterRepository
	CharaRepository domainCharacter.CharacterRepository
}

type GachaDrawRequest struct {
	Times int `json: "times"`
}

type GachaDrawResponse struct {
	Results []GachaResult `json: "results"`
}

type GachaResult struct {
	CharacterID int    `json: "characterID"`
	Name        string `json: "name"`
}

type CharacterListResponse struct {
	Characters []domainUserCharacter.UserCharacter
}

func (r GachaApplication) Gacha(times, uid, count int) (GachaDrawResponse, error) {

	if times < 1 {
		return GachaDrawResponse{}, &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalied times",
		}
	}

	charaList := GachaDrawResponse{}

	for i := 0; i < times; i++ {
		chara, err := r.doGacha(count)
		if err != nil {
			return GachaDrawResponse{}, err
		}

		charaList.Results = append(charaList.Results, *chara)

		var userChara domainUserCharacter.UserCharacter
		userChara.CharacterID = chara.CharacterID
		userChara.Name = chara.Name
		userChara.UserCharacterID = uid
		err = r.Repository.Create(userChara)
		if err != nil {
			return GachaDrawResponse{}, err
		}
	}

	return charaList, nil
}

func (r GachaApplication) doGacha(count int) (*GachaResult, error) {
	rand.Seed(time.Now().UnixNano())
	id := rand.Intn(count)
	ok := true
	if id == 0 {
		ok = false
	}
	if !ok {
		id = rand.Intn(count)
		if id != 0 {
			ok = true
		}
	}
	result := new(GachaResult)
	result.CharacterID = id
	name, err := r.Repository.Find(id)
	if err != nil {
		return &GachaResult{}, err
	}
	result.Name = name
	return result, nil
}

func (r GachaApplication) FindAll(id int) ([]domainUserCharacter.UserCharacter, error) {
	return r.Repository.FindAll(id)
}
