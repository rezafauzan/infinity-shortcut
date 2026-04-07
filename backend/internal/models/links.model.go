package models

import "time"

type Links struct {
	Id          int       `json:"id" db:"id"`
	UserId      int       `json:"user_id" db:"user_id"`
	OriginalUrl string    `json:"original_url" db:"original_url"`
	ShortUrl    string    `json:"short_url" db:"short_url"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at" db:"deleted_at"`
}
