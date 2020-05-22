package "models"

type GachaDrawRequest struct {
	Time int `json: "time"`
}

type GachaDrawResponce struct {
	Result GachaResult
}

type GachaResult struct {
	CharacterID string `json: "characterID"`
	Name        string `json: "name"`
}


