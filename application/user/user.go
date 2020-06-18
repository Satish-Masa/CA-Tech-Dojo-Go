package user

import domainUser "github.com/Satish-Masa/CA-Tech-Dojo-Go/domain/user"

type UserApplication struct {
	Repository domainUser.UserRepository
}

type UserCreatRequest struct {
	Name string `json: "name"`
}

type UserUpdateRequest struct {
	Name string `json: "name"`
}

type UserGetResponce struct {
	Name string `json: "name"`
}

type UserCreatResponse struct {
	Token string `json: "token"`
}

func (a UserApplication) SaveUser(u *domainUser.User) error {
	return a.Repository.Save(u)
}

func (a UserApplication) FindUser(uid int) (UserGetResponce, error) {
	user, err := a.Repository.Find(&domainUser.User{ID: uid})
	if err != nil {
		return UserGetResponce{}, err
	}
	name := user.Name
	resp := UserGetResponce{Name: name}
	return resp, nil
}

func (a UserApplication) UpdateUser(name string, id int) error {
	return a.Repository.Update(name, id)
}

/* func (a UserApplication) GetList(u domain.User) (CharacterListResponse, error) {
	return infrastructure.FindChara(u.Token)
} */
