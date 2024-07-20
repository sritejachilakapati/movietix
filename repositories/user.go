package repositories

import "movietix/models"

type UserRepository interface {
	GetUserByID(id int) (*models.User, error)
}

type pgUserRepository struct{}

func (r *pgUserRepository) GetUserByID(id int) (*models.User, error) {
	return &models.User{
		ID:   id,
		Name: "John Doe",
		Age:  30,
	}, nil
}

func NewUserRepository() UserRepository {
	return &pgUserRepository{}
}
