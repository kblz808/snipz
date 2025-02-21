package services

import (
	"context"
	"snipz/internal/storage/repository"
)

type UserService struct {
	repo repository.UserRepository
	// TODO: use cache
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo}
}

func (service *UserService) Register(ctx context.Context, user *repository.User)
