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

func (u LinkRepository) CreateNewLink(newLink models.Links) (models.Links, error) {
	sql := "INSERT INTO links (user_id, original_url, short_url, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id, user_id, original_url, short_url, created_at, updated_at, deleted_at"
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

func (u LinkRepository) GetLinkById(linkId int) (models.Links, error) {
	sql := "SELECT id, user_id, original_url, short_url, created_at, updated_at, deleted_at FROM links WHERE id = $1"

	rows, err := u.db.Query(context.Background(), sql, linkId)
	if err != nil {
		return models.Links{}, errors.New("Failed to get links! : " + err.Error())
	}

	link, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Links])
	if err != nil {
		return models.Links{}, errors.New("Links fetched but returning error! : " + err.Error())
	}

	return link, nil
}

func (u LinkRepository) GetAllLinksByUserId(userId int) ([]models.Links, error) {
	sql := "SELECT id, user_id, original_url, short_url, created_at, updated_at, deleted_at FROM links WHERE user_id = $1"

	rows, err := u.db.Query(context.Background(), sql, userId)
	if err != nil {
		return []models.Links{}, errors.New("Failed to get links! : " + err.Error())
	}

	links, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Links])
	if err != nil {
		return []models.Links{}, errors.New("Links fetched but returning error! : " + err.Error())
	}

	return links, nil
}

func (u LinkRepository) DeleteLinkById(id int) (models.Links, error) {
	sql := "UPDATE links SET deleted_at = $1 WHERE id = $2 RETURNING id, user_id, original_url, short_url, created_at, updated_at, deleted_at"

	rows, err := u.db.Query(context.Background(), sql, time.Now(), id)
	if err != nil {
		return models.Links{}, errors.New("Failed to delete link! : " + err.Error())
	}

	deletedLink, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Links])
	if err != nil {
		return models.Links{}, errors.New("Link deleted but returning error! : " + err.Error())
	}

	return deletedLink, nil
}