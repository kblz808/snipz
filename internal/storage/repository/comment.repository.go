package repository

import (
	"context"
	"errors"
	"snipz/internal/storage"
	"time"
)

type Comment struct {
	ID int64

	SnippetID int64

	UserID int64

	Content string

	CreatedAt time.Time
	UpdatedAt time.Time
}

type CommentRepository struct {
	db *storage.DB
}

func NewCommentRepository(db *storage.DB) *CommentRepository {
	return &CommentRepository{db}
}

func (repo *CommentRepository) CreateComment(ctx context.Context, comment *Comment, userId int64) (*Comment, error) {
	query := repo.db.QueryBuilder.Insert("comments").
		Columns("snippet_id", "user_id", "content").
		Values(comment.SnippetID, userId, comment.Content).
		Suffix("RETURNING *")

	sql, args, err := query.ToSql()
	if err != nil {
		if errCode := repo.db.ErrorCode(err); errCode == "23505" {
			return nil, errors.New("conflicting data")
		}
		return nil, err
	}

	err = repo.db.QueryRow(ctx, sql, args...).Scan(
		&comment.ID,
		&comment.SnippetID,
		&comment.UserID,
		&comment.Content,
		&comment.CreatedAt,
		&comment.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return comment, nil
}

func (repo *CommentRepository) GetSnippetComments(ctx context.Context, snippetId int64) ([]Comment, error) {
	var comment Comment
	var comments []Comment

	query := repo.db.QueryBuilder.Select("*").
		From("comments").
		OrderBy("date_created")

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := repo.db.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(
			&comment.ID,
			&comment.SnippetID,
			&comment.UserID,
			&comment.Content,
			&comment.CreatedAt,
			&comment.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		comments = append(comments, comment)
	}

	return comments, nil
}
