package repository

import (
	"context"
	"errors"
	"snipz/internal/storage"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

type Visibility string

const (
	Public    Visibility = "public"
	Private   Visibility = "private"
	Protected Visibility = "protected"
)

type Snippet struct {
	ID int64

	Title    string
	Content  string
	Language string

	UserID int64

	ExpiresAt time.Time

	Visibility Visibility

	IsEncrypted  bool
	PasswordHash string

	ViewCount int64
	Tags      []string

	StarsCount int64

	CreatedAt time.Time
	UpdatedAt time.Time
}

type SnippetRepository struct {
	db *storage.DB
}

func NewSnippetRepository(db *storage.DB) *SnippetRepository {
	return &SnippetRepository{db}
}

func (repo *SnippetRepository) CreateSnippet(ctx context.Context, snippet *Snippet) (*Snippet, error) {
	query := repo.db.QueryBuilder.Insert("snippets").
		Columns("title", "content", "language", "user_id", "expires_at", "visibility", "is_encrypted", "password_hash").
		Values(snippet.Title, snippet.Content, snippet.Language, snippet.UserID, snippet.ExpiresAt, snippet.Visibility, snippet.IsEncrypted, snippet.PasswordHash).
		Suffix("RETURNING *")

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	err = repo.db.QueryRow(ctx, sql, args...).Scan(
		&snippet.ID,
		&snippet.Title,
		&snippet.Content,
		&snippet.Language,
		&snippet.UserID,
		&snippet.ExpiresAt,
		&snippet.Visibility,
		&snippet.IsEncrypted,
		&snippet.PasswordHash,
		&snippet.ViewCount,
		&snippet.Tags,
		&snippet.StarsCount,
		&snippet.CreatedAt,
	)
	if err != nil {
		if errCode := repo.db.ErrorCode(err); errCode == "23505" {
			// TODO: replace with domain specific error
			return nil, errors.New("conflicting data")
		}
		return nil, err
	}

	return snippet, nil
}

func (repo *SnippetRepository) GetSnippetByID(ctx context.Context, id uint64) (*Snippet, error) {
	var snippet Snippet

	query := repo.db.QueryBuilder.Select("*").
		From("snippets").
		Where(sq.Eq{"id": id}).
		Limit(1)

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	err = repo.db.QueryRow(ctx, sql, args...).Scan(
		&snippet.ID,
		&snippet.Title,
		&snippet.Content,
		&snippet.Language,
		&snippet.UserID,
		&snippet.ExpiresAt,
		&snippet.Visibility,
		&snippet.IsEncrypted,
		&snippet.PasswordHash,
		&snippet.ViewCount,
		&snippet.Tags,
		&snippet.StarsCount,
		&snippet.CreatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			// TODO: replace with domain specific error
			return nil, errors.New("data not found")
		}
		return nil, err
	}

	return &snippet, nil
}

func (repo *SnippetRepository) ListSnippets(ctx context.Context, title, language string, skip, limit uint64) ([]Snippet, error) {
	var snippet Snippet
	var snippets []Snippet

	query := repo.db.QueryBuilder.Select("*").
		From("snippets").
		OrderBy("id").
		Limit(limit).
		Offset((skip - 1) * limit)

	if language != "" {
		query = query.Where(sq.Eq{"language": language})
	}

	if title != "" {
		query = query.Where(sq.ILike{"title": "%" + title + "%"})
	}

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := repo.db.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err := rows.Scan(
			&snippet.ID,
			&snippet.Title,
			&snippet.Content,
			&snippet.Language,
			&snippet.UserID,
			&snippet.ExpiresAt,
			&snippet.Visibility,
			&snippet.IsEncrypted,
			&snippet.PasswordHash,
			&snippet.ViewCount,
			&snippet.Tags,
			&snippet.StarsCount,
			&snippet.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		snippets = append(snippets, snippet)
	}

	return snippets, nil
}

func (repo *SnippetRepository) DeleteSnippet(ctx context.Context, id uint64) error {
	query := repo.db.QueryBuilder.Delete("snippets").
		Where(sq.Eq{"id": id})

	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}

	_, err = repo.db.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}

	return nil
}
