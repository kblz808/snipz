package services

import (
	"context"
	"snipz/internal/storage/repository"
)

type CommentService struct {
	repo repository.CommentRepository
}

func NewCommentService(repo repository.CommentRepository) *CommentService {
	return &CommentService{repo}
}

func (service *CommentService) PostComment(ctx context.Context, comment *repository.Comment, userID int64) (*repository.Comment, error) {
	comment, err := service.repo.CreateComment(ctx, comment, userID)
	if err != nil {
		return nil, err
	}

	return comment, nil
}

func (service *CommentService) GetCommentsForSnippet(ctx context.Context, snippetID int64) ([]*repository.Comment, error) {
	comments, err := service.GetCommentsForSnippet(ctx, snippetID)
	if err != nil {
		return nil, err
	}

	return comments, nil
}
