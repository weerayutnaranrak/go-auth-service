package service

type NewUserRequest struct {
	Name string `json:"name"`
}

type UserResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type UserService interface {
	NewUser(NewUserRequest) (*UserResponse, error)
	GetUsers(int) ([]UserResponse, error)
}
