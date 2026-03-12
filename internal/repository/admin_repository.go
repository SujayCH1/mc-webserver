package repository

import (
	"context"
	"database/sql"
	"time"
)

type AdminUser struct {
	ID           int
	Username     string
	PasswordHash string
	Role         string
	CreatedAt    time.Time
}

type AdminRepository struct {
	DB *sql.DB
}

func NewAdminRepository(db *sql.DB) *AdminRepository {
	return &AdminRepository{DB: db}
}

func (r *AdminRepository) GetByUsername(ctx context.Context, username string) (*AdminUser, error) {

	query := `
	SELECT id, username, password_hash, role, created_at
	FROM admin_users
	WHERE username=$1
	`

	row := r.DB.QueryRowContext(ctx, query, username)

	var a AdminUser

	err := row.Scan(
		&a.ID,
		&a.Username,
		&a.PasswordHash,
		&a.Role,
		&a.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &a, nil
}
