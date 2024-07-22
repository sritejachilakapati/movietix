package services

import (
	"context"
	"movietix/models"
	"movietix/repositories"
)

type UserService interface {
	GetUserByID(ctx context.Context, id string) (*models.User, error)
	CreateUser(ctx context.Context, user *models.User) error
}

type userService struct {
	userRepo repositories.UserRepository
}

func (s *userService) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	return s.userRepo.GetByID(ctx, id)
}

func (s *userService) CreateUser(ctx context.Context, user *models.User) error {
	return s.userRepo.Create(ctx, user)
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{userRepo}
}
