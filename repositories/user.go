package repositories

import (
	"movietix/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByID(id string) (*models.User, error)
	CreateUser(user *models.User) error
}

type pgUserRepository struct {
	db *gorm.DB
}

func (r *pgUserRepository) GetUserByID(id string) (*models.User, error) {
	var user models.User
	result := r.db.First(&user, "id = ?", id)
	return &user, result.Error
}

func (r *pgUserRepository) CreateUser(user *models.User) error {
	result := r.db.Create(user)
	return result.Error
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &pgUserRepository{db}
}
