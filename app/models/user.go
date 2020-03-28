package models

type User struct {
	Name  string `json:"name"`
	Token string `gorm: "primary_key"`
}

type UserCreatReqest struct {
	Name string
}

type UserCreatResponse struct {
	Token string
}

type UserGetResponse struct {
	Name string
}

type UserUpdateRequest struct {
	Name string
}
