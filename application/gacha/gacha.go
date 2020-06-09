package gacha

import (
	"math/rand"
	"time"

	"github.com/Satish-Masa/CA-Tech-Dojo-Go/domain"
	"github.com/Satish-Masa/CA-Tech-Dojo-Go/domain/repository"
	"github.com/Satish-Masa/CA-Tech-Dojo-Go/infrastructure"
)

type GachaApplication struct {
	Repository repository.CharacterRepository
}

type GachaDrawRequest struct {
	Times int `json: "times"`
}

type GachaDrawResponse struct {
	Results []GachaResult
}

type GachaResult struct {
	CharacterID int    `json: "characterID"`
	Name        string `json: "name"`
}

// DoGacha()で回数のリクエストを受け取り、結果のレスポンスを返す。
func (r GachaApplication) DoGacha(g *GachaDrawRequest) GachaDrawResponse {
	times := g.Times
	result := r.run(times)
	var resp GachaDrawResponse
	resp.Results = result

	return resp
}

// ガチャの結果を決める抽選するメソッド
func (r GachaApplication) run(times int) []GachaResult {
	count, _ := infrastructure.CharaCount()
	result := make([]GachaResult, times)

	for i := 0; i < times; i++ {
		var res GachaResult
		rand.Seed(time.Now().UnixNano())
		res.CharacterID = rand.Intn(count)
		res.Name, _ = infrastructure.FindChara(res.CharacterID)
		err := r.update(res)
		if err != nil {
			panic(err)
		}
	}

	return result
}

// データベースにガチャの結果を保存するメソッド
func (r GachaApplication) update(g GachaResult) error {
	var chara domain.Character
	chara.CharacterID = g.CharacterID
	// ガチャをしたUserのTokenをUserCharacterIDに入れる
	// chara.UserCharacterID = token
	chara.Name = g.Name

	err := infrastructure.UpdateChar(chara)
	if err != nil {
		return err
	}
	return nil
}
