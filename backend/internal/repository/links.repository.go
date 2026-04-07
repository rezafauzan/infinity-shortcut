package repository

import (
	"context"
	"errors"
	"snowfoxinfinity/infinity-shortcut/internal/models"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type LinkRepository struct {
	db *pgxpool.Pool
}

func NewLinkRepository(db *pgxpool.Pool) (*LinkRepository, error) {
	return &LinkRepository{
		db: db,
	}, nil
}

func (u UserRepository) CreateNewLink(newLink models.Links) (models.Links, error) {
	sql := "INSERT INTO users (user_id, original_url, short_url, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id, user_id, original_url, short_url, created_at, updated_at"
	rows, err := u.db.Query(context.Background(), sql, newLink.UserId, newLink.OriginalUrl, newLink.ShortUrl, time.Now(), time.Now())
	if err != nil {
		return models.Links{}, errors.New("Failed to create new link! : " + err.Error())
	}

	createdLink, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Links])
	if err != nil {
		return models.Links{}, errors.New("Link created but returning error! : " + err.Error())
	}

	return createdLink, nil
}