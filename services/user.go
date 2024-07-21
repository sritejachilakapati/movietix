package services

import (
	"movietix/models"
	"movietix/repositories"
)

type UserService interface {
	GetUserByID(id string) (*models.User, error)
	CreateUser(user *models.User) error
}

type userService struct {
	userRepo repositories.UserRepository
}

func (s *userService) GetUserByID(id string) (*models.User, error) {
	return s.userRepo.GetUserByID(id)
}

func (s *userService) CreateUser(user *models.User) error {
	return s.userRepo.CreateUser(user)
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{userRepo}
}
