package user

type UserRepository interface {
	Save(*User) error
	Find(*User) (User, error)
	Update(string, int) error
}
