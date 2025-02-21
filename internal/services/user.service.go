package services

import (
	"context"
	"errors"
	"snipz/internal/storage/repository"
	"snipz/internal/utils"
)

type UserService struct {
	repo repository.UserRepository
	// TODO: use cache
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo}
}

func (service *UserService) Register(ctx context.Context, user *repository.User) (*repository.User, error) {
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		// TODO: replace with domain specific error
		return nil, errors.New("internal error")
	}
	user.Password = hashedPassword

	user, err = service.repo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (service *UserService) GetUser(ctx context.Context, id uint64) (*repository.User, error) {
	var user *repository.User

	user, err := service.repo.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
