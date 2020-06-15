package gacha

import (
	"math/rand"
	"net/http"
	"time"

	domainCharacter "github.com/Satish-Masa/CA-Tech-Dojo-Go/domain/character"
	"github.com/Satish-Masa/CA-Tech-Dojo-Go/interfaces"
	"github.com/labstack/echo/v4"
)

type GachaApplication struct {
	Repository domainCharacter.CharacterRepository
}

type GachaDrawResponse struct {
	Results []GachaResult
}

type GachaResult struct {
	CharacterID int    `json: "characterID"`
	Name        string `json: "name"`
}

func (r GachaApplication) Gacha(g *interfaces.GachaDrawRequest, uid int) (result GachaDrawResponse, err error) {
	if g.Times < 1 {
		return GachaDrawResponse{}, &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalied time",
		}
	}

	count, err := r.Repository.CharaCount()
	if err != nil {
		return GachaDrawResponse{}, err
	}

	var res []GachaResult

	if g.Times == 1 {
		res, err := r.gachaOneTime(count)
		if err != nil {
			return GachaDrawResponse{}, err
		}
		character := domainCharacter.NewCharacter(uid, res[0].CharacterID, res[0].Name)
		err = r.Repository.CreateChara(*character)
		if err != nil {
			return GachaDrawResponse{}, err
		}
	} else {
		res, err := r.gachaManyTime(count, g.Times)
		if err != nil {
			return GachaDrawResponse{}, err
		}

		for i := 0; i < g.Times; i++ {
			character := domainCharacter.NewCharacter(uid, res[i].CharacterID, res[i].Name)
			err := r.Repository.CreateChara(*character)
			if err != nil {
				return GachaDrawResponse{}, err
			}
		}
	}

	result.Results = res

	return result, nil
}

func (r GachaApplication) gachaOneTime(count int) (result []GachaResult, err error) {
	result[0], err = r.doGacha(count)
	if err != nil {
		return []GachaResult{}, err
	}

	return result, nil
}

func (r GachaApplication) gachaManyTime(count, times int) ([]GachaResult, error) {
	result := make([]GachaResult, times)
	for i := 0; i < times; i++ {
		chara, err := r.doGacha(count)
		if err != nil {
			return []GachaResult{}, err
		}
		result[i] = chara
	}

	return result, nil
}

func (r GachaApplication) doGacha(count int) (result GachaResult, err error) {
	rand.Seed(time.Now().UnixNano())
	result.CharacterID = rand.Intn(count)
	result.Name, err = r.Repository.FindChara(result.CharacterID)
	if err != nil {
		return GachaResult{}, err
	}

	return result, nil
}
