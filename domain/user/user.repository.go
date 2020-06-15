package user

type UserRepository interface {
	Save(*User) error
	Find(int) (string, error)
	Update(string, int) error
}
