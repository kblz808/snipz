package repository

import "time"

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
