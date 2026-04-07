package models

type Links struct {
	Id          string `json:"id" db:"id"`
	UserId      string `json:"user_id" db:"user_id"`
	OriginalUrl string `json:"original_url" db:"original_url"`
	ShortUrl    string `json:"short_url" db:"short_url"`
	CreatedAt   string `json:"created_at" db:"created_at"`
	UpdatedAt   string `json:"updated_at" db:"updated_at"`
	DeletedAt   string `json:"deleted_at" db:"deleted_at"`
}
