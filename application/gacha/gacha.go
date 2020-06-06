package gacha

import (
	"crypto/rand"
	"math/big"

	"github.com/Satish-Masa/CA-Tech-Dojo-Go/domain"
	"github.com/Satish-Masa/CA-Tech-Dojo-Go/infrastructure"
)

type GachaDrawRequest struct {
	Time int `json: "time"`
}

type GachaDrawResponse struct {
	Results []GachaResult
}

type GachaResult struct {
	CharacterID int    `json: "characterID"`
	Name        string `json: "name"`
}

func DoGacha(g *GachaDrawRequest) *GachaDrawResponse {
	time := g.Time
	result := run(time)
	var resp GachaDrawResponse
	resp = result

	return &resp
}

func run(int time) []GachaResult {
	count := infrastructure.CharaCount()
	result := make([]GachaResult, time)

	for i := 0; i < time; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(count))
		if err != nil {
			panic(err)
		}
		result[i].CharacterID = n
		name := infrastructure.FindChara(n)
		result[i].Name = name

		err := update(result)
		if err != nil {
			panic(err)
		}
	}

	return result
}

func update(r *GachaResult) error {
	var chara domain.Character
	chara.CharacterID = r.CharacterID
	// ガチャをしたUserのTokenをUserCharacterIDに入れる
	// chara.UserCharacterID = token
	chara.Name = r.Name

	err := UpdateChar(chara)
	if err != nil {
		return err
	}
	return nil
}
