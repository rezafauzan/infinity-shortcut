package dto

type CreateNewLinkDTO struct {
	UserId      int `json:"user_id" binding:"required"`
	OriginalUrl string `json:"original_url" binding:"required"`
	ShortUrl    string `json:"short_url" binding:"required"`
}

type CreateNewLinkResponseDTO struct {
	Id          int `json:"id"`
	UserId      int `json:"user_id"`
	OriginalUrl string `json:"original_url"`
	ShortUrl    string `json:"short_url"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
