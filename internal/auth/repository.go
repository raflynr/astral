package auth

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type AuthRepository interface {
	Register(ctx context.Context, register Register) error
	Login(ctx context.Context, login Login) (*User, error)
}

type authRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) AuthRepository {
	return &authRepository{
		db: db,
	}
}

func (ar *authRepository) Register(ctx context.Context, register Register) error {
	query := "INSERT INTO users (id, email, password, username) VALUES ($1, $2, $3, $4)"
	_, err := ar.db.ExecContext(ctx, query, register.ID, register.Email, register.Password, register.Username)
	if err != nil {
		return err
	}

	return nil
}

func (ar *authRepository) Login(ctx context.Context, login Login) (*User, error) {
	var user User

	query := `SELECT id, email, password FROM users WHERE email = $1`
	row := ar.db.QueryRowContext(ctx, query, login.Email)

	err := row.Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("email tidak ditemukan")
		}
		return nil, fmt.Errorf("gagal mengambil data user: %w", err)
	}

	return &user, nil
}
