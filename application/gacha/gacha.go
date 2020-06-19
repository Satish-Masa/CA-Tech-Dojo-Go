package gacha

import (
	"fmt"
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

func (r GachaApplication) Gacha(times, uid, count int) (result GachaDrawResponse, err error) {
	if times < 1 {
		return GachaDrawResponse{}, &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalied times",
		}
	}

	charaList := make([]GachaResult, 0, times)
	for _, chara := range charaList {
		res, err := r.doGacha(count)
		if err != nil {
			return GachaDrawResponse{}, &echo.HTTPError{
				Code:    http.StatusInternalServerError,
				Message: "failed to do gacha",
			}
		}
		chara.CharacterID = res.CharacterID
		fmt.Printf("CharaID: %d\n", chara.CharacterID)
		chara.Name = res.Name
		charaList = append(charaList, chara)
	}

	for _, chara := range charaList {
		var userChara domainUserCharacter.UserCharacter
		userChara.CharacterID = chara.CharacterID
		userChara.Name = chara.Name
		userChara.UserCharacterID = uid
		err := r.Repository.Create(userChara)
		if err != nil {
			return GachaDrawResponse{}, err
		}
	}

	result.Results = charaList

	return result, nil
}

func (r GachaApplication) doGacha(count int) (result GachaResult, err error) {
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
	result.CharacterID = id
	chara, err := r.CharaRepository.Find(id)
	result.Name = chara.Name
	if err != nil {
		return GachaResult{}, err
	}
	return result, nil
}
