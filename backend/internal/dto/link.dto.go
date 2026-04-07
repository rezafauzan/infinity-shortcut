package dto

import "time"

type CreateNewLinkDTO struct {
	OriginalUrl string `json:"original_url" binding:"required"`
}

type CreateNewLinkResponseDTO struct {
	Id          int       `json:"id"`
	UserId      int       `json:"user_id"`
	OriginalUrl string    `json:"original_url"`
	ShortUrl    string    `json:"short_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type GetAllLinksResponseDTO struct {
	Id          int       `json:"id"`
	UserId      int       `json:"user_id"`
	OriginalUrl string    `json:"original_url"`
	ShortUrl    string    `json:"short_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}