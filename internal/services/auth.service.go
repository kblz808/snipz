package services

import (
	"context"
	"snipz/internal/storage/repository"
	"snipz/internal/utils"
)

type AuthService struct {
	userRepo repository.UserRepository
	ts       TokenService
}

func NewAuthService(repo repository.UserRepository, ts TokenService) *AuthService {
	return &AuthService{repo, ts}
}

func (as *AuthService) Login(ctx context.Context, email, password string) (string, error) {
	user, err := as.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	err = utils.ComparePassword(password, user.Password)
	if err != nil {
		return "", err
	}

	token, err := as.ts.CreateToken(*user)
	if err != nil {
		return "", err
	}

	return token, nil
}
