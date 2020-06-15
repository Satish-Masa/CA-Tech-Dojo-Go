package user

type UserRepository interface {
	Save(*User) error
	Find(int) (User, error)
	Update(string, int) error
}
