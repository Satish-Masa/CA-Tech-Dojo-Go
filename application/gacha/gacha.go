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
	Results []gachaResult
}

type gachaResult struct {
	CharacterID int    `json: "characterID"`
	Name        string `json: "name"`
}

func (r GachaApplication) Gacha(times, uid, count int) (result GachaDrawResponse, err error) {
	if times < 1 {
		return GachaDrawResponse{}, &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalied time",
		}
	}

	if err != nil {
		return GachaDrawResponse{}, err
	}

	var res []gachaResult

	if times == 1 {
		res, err := r.gachaOneTime(count)
		if err != nil {
			return GachaDrawResponse{}, err
		}
		character := domainUserCharacter.NewCharacter(uid, res[0].CharacterID, res[0].Name)
		err = r.Repository.Create(*character)
		if err != nil {
			return GachaDrawResponse{}, err
		}
	} else {
		res, err := r.gachaManyTime(count, times)
		if err != nil {
			return GachaDrawResponse{}, err
		}

		for i := 0; i < times; i++ {
			character := domainUserCharacter.NewCharacter(uid, res[i].CharacterID, res[i].Name)
			err := r.Repository.Create(*character)
			if err != nil {
				return GachaDrawResponse{}, err
			}
		}
	}

	result.Results = res

	return result, nil
}

func (r GachaApplication) gachaOneTime(count int) (result []gachaResult, err error) {
	result[0], err = r.doGacha(count)
	if err != nil {
		return []gachaResult{}, err
	}

	return result, nil
}

func (r GachaApplication) gachaManyTime(count, times int) ([]gachaResult, error) {
	result := make([]gachaResult, times)
	for i := 0; i < times; i++ {
		chara, err := r.doGacha(count)
		if err != nil {
			return []gachaResult{}, err
		}
		result[i] = chara
	}

	return result, nil
}

func (r GachaApplication) doGacha(count int) (result gachaResult, err error) {
	rand.Seed(time.Now().UnixNano())
	result.CharacterID = rand.Intn(count)
	chara, err := r.CharaRepository.Find(count)
	result.Name = chara.Name
	if err != nil {
		return gachaResult{}, err
	}
	return result, nil
}
