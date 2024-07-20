package services

import (
	"movietix/models"
	"movietix/repositories"
)

type UserService interface {
	GetUserByID(id int) (*models.User, error)
}

type userService struct {
	userRepo repositories.UserRepository
}

func (s *userService) GetUserByID(id int) (*models.User, error) {
	return s.userRepo.GetUserByID(id)
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{userRepo}
}
