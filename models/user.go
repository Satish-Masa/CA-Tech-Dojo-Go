package models

type User struct {
	name  string
	token string
}

type UserCreatRequest struct {
	name string
}

type UserCreatResponse struct {
	token string
}

type UserGetResponce struct {
	name string
}

type UserUpdateRequest struct {
	name string
}
