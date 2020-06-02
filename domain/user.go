package domain

import "fmt"

type User struct {
	Name  string `json: "name"`
	Token string `json: "token"`
}

func NewUser(name, token string) (*User, error) {
	if name == "" {
		return nil, fmt.Errorf("invalid name")
	}

	return &User{
		Name:  name,
		Token: token,
	}, nil
}
