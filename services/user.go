package services

import (
	"context"
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
	return s.userRepo.GetByID(context.Background(), id)
}

func (s *userService) CreateUser(user *models.User) error {
	return s.userRepo.Create(context.Background(), user)
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{userRepo}
}
