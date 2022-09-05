package service

import (
	"auth-service/repository"

	"github.com/jinzhu/copier"
)

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return userService{repo: repo}
}

func (s userService) NewUser(user NewUserRequest) (*UserResponse, error) {
	newUser := repository.User{Name: user.Name}
	data, err := s.repo.CreateUser(newUser)
	if err != nil {
		println(err)
	}
	res := UserResponse{Name: data.Name, ID: int(data.ID)}
	return &res, nil
}

func (s userService) GetUsers(int) ([]UserResponse, error) {
	users, _ := s.repo.GetUsers()
	userRes := []UserResponse{}
	copier.Copy(&userRes, &users)

	return userRes, nil
}
