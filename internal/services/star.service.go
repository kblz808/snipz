package services

import (
	"context"
	"snipz/internal/storage/repository"
)

type StarService struct {
	repo repository.StarRepository
}

func NewStarService(repo repository.StarRepository) *StarService {
	return &StarService{repo}
}

func (service *StarService) StarSnippet(ctx context.Context, snippetID, userID int64) (*repository.Star, error) {
	star := &repository.Star{
		SnippetID: snippetID,
		UserID:    userID,
	}

	star, err := service.repo.CreateStar(ctx, star)
	if err != nil {
		return nil, err
	}

	return star, nil
}
