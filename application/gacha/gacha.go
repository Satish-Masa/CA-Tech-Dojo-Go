package gacha

import (
	"crypto/rand"
	"math/big"

	"github.com/Satish-Masa/CA-Tech-Dojo-Go/domain"
	"github.com/Satish-Masa/CA-Tech-Dojo-Go/domain/repository"
	"github.com/Satish-Masa/CA-Tech-Dojo-Go/infrastructure"
)

type GachaApplication struct {
	Repository repository.CharacterRepository
}

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

// DoGacha()で回数のリクエストを受け取り、結果のレスポンスを返す。
func (r *GachaApplication) DoGacha(g GachaDrawRequest) GachaDrawResponse {
	time := g.Time
	result := r.run(time)
	var resp GachaDrawResponse
	resp.Results = result

	return resp
}

// ガチャの結果を決める抽選するメソッド
func (r *GachaApplication) run(time int) []GachaResult {
	count, _ := infrastructure.CharaCount()
	result := make([]GachaResult, time)

	for i := 0; i < time; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(count))
		if err != nil {
			panic(err)
		}
		result[i].CharacterID = n
		name, _ := infrastructure.FindChara(n)
		result[i].Name = name

		err := r.update(result)
		if err != nil {
			panic(err)
		}
	}

	return result
}

// データベースにガチャの結果を保存するメソッド
func (r *GachaApplication) update(g GachaResult) error {
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
