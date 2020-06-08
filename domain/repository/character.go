package repository

import "github.com/Satish-Masa/CA-Tech-Dojo-Go/application/gacha"

type CharacterRepository interface {
	DoGacha(gacha.GachaDrawRequest) gacha.GachaDrawResponse
}
