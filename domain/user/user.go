package user

type User struct {
	ID   int    `json: "id" gorm: "praimaly_key"`
	Name string `json: "name"`
}

func NewUser(name string) *User {
	return &User{
		Name: name,
	}
}
