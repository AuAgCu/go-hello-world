package service

import (
	"app/models"
	"app/repository"
)

type UserService interface {
	CreateUser(lastName string, firstName string)
	GetUserById(id int) models.User
}

type UserServiceImpl struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &UserServiceImpl{userRepository}
}

func (this UserServiceImpl) CreateUser(lastName string, firstName string) {
	user := models.User{FIRST_NAME: firstName, LAST_NAME: lastName}
	this.userRepository.CreateUser(&user)
}

func (userService UserServiceImpl) GetUserById(id int) models.User {
	return userService.GetUserById(id)
}
