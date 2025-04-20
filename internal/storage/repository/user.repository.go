package repository

import (
	"context"
	"errors"
	"snipz/internal/storage"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

type User struct {
	ID int64

	Username string
	Password string

	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserRepository struct {
	db *storage.DB
}

func NewUserRepository(db *storage.DB) *UserRepository {
	return &UserRepository{db}
}

func (ur *UserRepository) CreateUser(ctx context.Context, user *User) (*User, error) {
	query := ur.db.QueryBuilder.Insert("users").
		Columns("username", "password").
		Values(user.Username, user.Password).
		Suffix("RETURNING *")

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	err = ur.db.QueryRow(ctx, sql, args...).Scan(
		&user.ID,
		&user.Username,
	)
	if err != nil {
		if errCode := ur.db.ErrorCode(err); errCode == "23505" {
			return nil, errors.New("conflicting data")
		}
		return nil, err
	}

	return user, nil
}

func (ur *UserRepository) GetUserByID(ctx context.Context, id uint64) (*User, error) {
	var user User

	query := ur.db.QueryBuilder.Select("*").
		From("users").
		Where(sq.Eq{"id": id}).
		Limit(1)

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	err = ur.db.QueryRow(ctx, sql, args...).Scan(
		&user.ID,
		&user.Username,
		&user.CreatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New("data not found")
		}
	}

	return &user, nil
}

func (ur *UserRepository) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	var user User

	query := ur.db.QueryBuilder.Select("*").
		From("users").
		Where(sq.Eq{"email": email}).
		Limit(1)

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	err = ur.db.QueryRow(ctx, sql, args...).Scan(
		&user.ID,
		&user.Username,
		&user.CreatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New("data not found")
		}
	}
	return &user, nil
}
