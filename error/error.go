package error

import "errors"

var (
	// infra layer
	// character.go
	ErrCount       = errors.New("failed to count characters")
	ErrCreateChara = errors.New("failed to save the character")

	// gacha.go
	ErrCreateUserChara = errors.New("failed to create the userCharacter")
	ErrFindChara       = errors.New("failed to find characters")
	ErrFindAll         = errors.New("failed to find your characters")

	// user.go
	ErrCreateUser = errors.New("failed to save the user")
	ErrFindUser   = errors.New("failed to find the user")
	ErrUpdateUser = errors.New("failed to update the user")
)
