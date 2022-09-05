package repository

type User struct {
	ID       uint `gorm:"primaryKey"`
	Name     string
	UserName string
	PassWord string
}

type UserRepository interface {
	GetUsers() ([]User, error)
	CreateUser(user User) (*User, error)
	SignIn(userName string, passWord string) (*User, error)
}
