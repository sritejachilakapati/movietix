package repositories

import (
	"context"
	"fmt"
	"strings"

	"github.com/sritejachilakapati/movietix/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	GetByID(ctx context.Context, id string) (*models.User, error)
	GetAll(ctx context.Context) ([]*models.User, error)
	Update(ctx context.Context, id string, updatePayload map[string]interface{}) error
	Delete(ctx context.Context, id string) error
	GetByQuery(ctx context.Context, queryParams map[string]interface{}) ([]*models.User, error)
}

type userRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, user *models.User) error {
	query := `INSERT INTO users (id, name, email, password, is_active, role, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := r.db.Exec(ctx, query, user.ID, user.Name, user.Email, user.Password, user.IsActive, user.Role, user.CreatedAt, user.UpdatedAt)
	return err
}

func (r *userRepository) GetByID(ctx context.Context, id string) (*models.User, error) {
	query := `SELECT id, name, email, is_active, role, created_at, updated_at FROM users WHERE id = $1`
	row := r.db.QueryRow(ctx, query, id)

	var user models.User
	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.IsActive,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetAll(ctx context.Context) ([]*models.User, error) {
	query := `SELECT id, name, email, is_active, role, created_at, updated_at FROM users`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.IsActive,
			&user.Role,
			&user.CreatedAt,
			&user.UpdatedAt,
		); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, rows.Err()
}

func (r *userRepository) Update(ctx context.Context, id string, updatePayload map[string]interface{}) error {
	var setQuery []string
	var args []interface{}
	var i = 1

	for key, value := range updatePayload {
		setQuery = append(setQuery, fmt.Sprintf("%s = $%d", key, i))
		args = append(args, value)
		i++
	}
	args = append(args, id)

	query := `UPDATE users SET`
	if len(setQuery) > 0 {
		query += " " + strings.Join(setQuery, ", ")
	}

	query += fmt.Sprintf(" WHERE id = $%d", i)

	_, err := r.db.Exec(ctx, query, args...)
	return err
}

func (r *userRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.Exec(ctx, query, id)
	return err
}

func (r *userRepository) GetByQuery(ctx context.Context, queryParams map[string]interface{}) ([]*models.User, error) {
	var conditions []string
	var args []interface{}
	var i = 1

	for key, value := range queryParams {
		conditions = append(conditions, fmt.Sprintf("%s = $%d", key, i))
		args = append(args, value)
		i++
	}

	query := `SELECT id, name, email, is_active, role, created_at, updated_at FROM users`
	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.IsActive,
			&user.Role,
			&user.CreatedAt,
			&user.UpdatedAt,
		); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, rows.Err()
}
