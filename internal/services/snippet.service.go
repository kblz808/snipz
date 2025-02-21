package services

import (
	"context"
	"snipz/internal/storage/repository"
)

type SnippetService struct {
	repo repository.SnippetRepository
}

func NewSnippetService(repo repository.SnippetRepository) *SnippetService {
	return &SnippetService{repo}
}

func (service *SnippetService) CreateSnippet(ctx context.Context, snippet *repository.Snippet) (*repository.Snippet, error) {
	snippet, err := service.repo.CreateSnippet(ctx, snippet)
	if err != nil {
		return nil, err
	}

	return snippet, nil
}

func (service *SnippetService) GetSnippet(ctx context.Context, id uint64) (*repository.Snippet, error) {
	var snippet *repository.Snippet

	snippet, err := service.repo.GetSnippetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return snippet, nil
}

func (service *SnippetService) GetAllSnippets(ctx context.Context, skip, limit uint64) ([]repository.Snippet, error) {
	snippets, err := service.repo.ListSnippets(ctx, "", "", skip, limit)
	if err != nil {
		return nil, err
	}

	return snippets, nil
}

func (service *SnippetService) SearchSnippet(ctx context.Context, title, langugage string, skip, limit uint64) ([]repository.Snippet, error) {
	snippets, err := service.repo.ListSnippets(ctx, title, langugage, skip, limit)
	return snippets, err
}
