package repository

import (
	"context"
	"snipz/internal/storage"
	"time"
)

type Star struct {
	ID int64

	SnippetID int64
	UserID    int64

	CreatedAt time.Time
}

type StarRepository struct {
	db *storage.DB
}

func NewStarRepository(db *storage.DB) *StarRepository {
	return &StarRepository{db}
}

func (repo *StarRepository) CreateStar(ctx context.Context, star *Star) (*Star, error) {
	query := repo.db.QueryBuilder.Insert("stars").
		Columns("user_id", "snippter_id").
		Values(star.UserID, star.SnippetID).
		Suffix("RETURNING *")

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	err = repo.db.QueryRow(ctx, sql, args...).Scan(
		&star.ID,
		&star.UserID,
		&star.SnippetID,
		&star.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return star, nil
}
