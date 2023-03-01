package repository

import (
	"auth-service/config"
	"fmt"

	"gorm.io/gorm"
)

type userRepositoryDB struct {
	db *gorm.DB
}

func NewUserRepositoryDB() UserRepository {
	return userRepositoryDB{db: config.GetDB()}
}

func (r userRepositoryDB) GetUsers() ([]User, error) {
	users := []User{}
	tx := r.db.Find(&users)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return nil, tx.Error
	}
	return users, nil
}

func (r userRepositoryDB) CreateUser(user User) (*User, error) {
	newUser := User{Name: user.Name}
	tx := r.db.Create(&newUser)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return nil, tx.Error
	}
	return &newUser, nil
}

func (r userRepositoryDB) SignIn(userName string, passWord string) (*User, error) {
	return nil, nil
}
