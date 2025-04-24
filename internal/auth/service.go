package auth

import (
	"context"
	"fmt"

	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"github.com/raflynr/astral/helper"
)

type AuthService interface {
	Register(ctx context.Context, register Register) error
	Login(ctx context.Context, login Login) (string, error)
}

type authService struct {
	authRepository AuthRepository
	Validate       *validator.Validate
}

func NewAuthService(authRepository AuthRepository, validate *validator.Validate) AuthService {
	return &authService{
		authRepository: authRepository,
		Validate:       validate,
	}
}

func (as *authService) Register(ctx context.Context, register Register) error {
	if err := as.Validate.Struct(register); err != nil {
		return helper.CustomMessageValidator(err)
	}

	id := uuid.NewString()
	register.ID = id

	hash, err := helper.HashPassword(register.Password)
	if err != nil {
		return fmt.Errorf("error when hashing password: %w", err)
	}
	register.Password = hash

	if err := as.authRepository.Register(ctx, register); err != nil {
		return fmt.Errorf("email already registered")
	}

	return nil
}

func (as *authService) Login(ctx context.Context, login Login) (string, error) {
	user, err := as.authRepository.Login(ctx, login)
	if err != nil {
		return "", fmt.Errorf("account not found")
	}

	if !helper.CheckPasswordHash(user.Password, login.Password) {
		return "", fmt.Errorf("email or password is incorrect")
	}

	token, err := helper.GenerateJWT(user.Email, user.Username)
	if err != nil {
		return "", err
	}

	return token, nil
}
