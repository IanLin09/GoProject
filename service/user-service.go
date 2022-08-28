package service

import (
	"goTest/models"
	"goTest/repository"
)

type UserService interface {
	FindUser(map[string]string) (models.User, error)
}

type userService struct {
	repository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{
		repository: userRepository,
	}
}

func (service *userService) FindUser(inputData map[string]string) (models.User, error) {
	user, err := service.repository.FindUser(inputData)

	return user, err
}
