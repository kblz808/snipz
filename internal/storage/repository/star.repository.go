package repository

import "time"

type Start struct {
	ID int64

	SnippetID int64
	User      int64

	CreatedAt time.Time
}
