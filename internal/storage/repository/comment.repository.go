package repository

import "time"

type Comment struct {
	ID int64

	SnippetID int64

	UserID int64

	Content string

	CreatedAt time.Time
	UpdatedAt time.Time
}
