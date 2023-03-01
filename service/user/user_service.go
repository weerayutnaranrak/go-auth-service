package service

import (
	userRepository "auth-service/repository/user"

	"github.com/jinzhu/copier"
)

type userService struct {
	repo userRepository.UserRepository
}

func NewUserService(repo userRepository.UserRepository) UserService {
	return userService{repo: repo}
}

func (s userService) NewUser(user NewUserRequest) (*UserResponse, error) {
	newUser := userRepository.User{Name: user.Name}
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
