package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"snowfoxinfinity/infinity-shortcut/internal/models"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type LinkRepository struct {
	db  *pgxpool.Pool
	rdb *redis.Client
}

func NewLinkRepository(db *pgxpool.Pool, rdb *redis.Client) (*LinkRepository, error) {
	return &LinkRepository{
		db:  db,
		rdb: rdb,
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

	if u.rdb != nil {
		cacheKey := fmt.Sprintf("links:%d", newLink.UserId)
		err := u.rdb.Del(context.Background(), cacheKey).Err()
		if err != nil {
			fmt.Println("REDIS DEL ERROR:", err)
		} else {
			fmt.Println("[CACHE INVALIDATED]", cacheKey)
		}
	}

	return createdLink, nil
}

func (u LinkRepository) GetLinkById(linkId int) (models.Links, error) {
	sql := "SELECT id, user_id, original_url, short_url, created_at, updated_at, deleted_at FROM links WHERE id = $1 AND deleted_at IS NULL"

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

func (u LinkRepository) GetLinkBySlug(slug string) (models.Links, error) {
	sql := "SELECT id, user_id, original_url, short_url, created_at, updated_at, deleted_at FROM links WHERE short_url = $1 AND deleted_at IS NULL"

	rows, err := u.db.Query(context.Background(), sql, slug)
	if err != nil {
		return models.Links{}, errors.New("Failed to get links! : " + err.Error())
	}

	link, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Links])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.Links{}, errors.New("Link not found or already deleted")
		}

		return models.Links{}, errors.New("Links fetched but returning error! : " + err.Error())
	}

	return link, nil
}

func (u LinkRepository) GetAllLinksByUserId(userId int) ([]models.Links, error) {
	cacheKey := fmt.Sprintf("links:%d", userId)

	if u.rdb != nil {
		valueCache, err := u.rdb.Get(context.Background(), cacheKey).Result()
		if err == nil {
			var links []models.Links
			if err := json.Unmarshal([]byte(valueCache), &links); err == nil {
				fmt.Println("[CACHE HIT]")
				return links, nil
			}
			fmt.Println("JSON UNMARSHAL ERROR:", err)
		} else if err != redis.Nil {
			fmt.Println("REDIS ERROR:", err)
		}
	}

	fmt.Println("[CACHE MISS]")

	sql := "SELECT id, user_id, original_url, short_url, created_at, updated_at, deleted_at FROM links WHERE user_id = $1 AND deleted_at IS NULL"

	rows, err := u.db.Query(context.Background(), sql, userId)
	if err != nil {
		return []models.Links{}, errors.New("Failed to get links! : " + err.Error())
	}

	links, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Links])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return []models.Links{}, errors.New("Link not found or already deleted")
		}

		return []models.Links{}, errors.New("Links fetched but returning error! : " + err.Error())
	}

	if u.rdb != nil && len(links) > 0 {
		data, err := json.Marshal(links)
		if err != nil {
			fmt.Println("JSON MARSHAL ERROR:", err)
		} else {
			err := u.rdb.Set(context.Background(), cacheKey, data, time.Hour).Err()
			if err != nil {
				fmt.Println("REDIS SET ERROR:", err)
			} else {
				fmt.Println("[CACHE SAVED]")
			}
		}
	}

	return links, nil
}

func (u LinkRepository) GetLinkByOriginalOrSlug(userId int, url string) ([]models.Links, error) {
	sql := `SELECT id, user_id, original_url, short_url, created_at, updated_at FROM links WHERE user_id = $1 AND deleted_at IS NULL AND (original_url ILIKE '%' || $2 || '%' OR short_url ILIKE '%' || $2 || '%')`

	rows, err := u.db.Query(context.Background(), sql, userId, url)
	if err != nil {
		return []models.Links{}, errors.New("Failed to get link: " + err.Error())
	}

	link, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Links])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return []models.Links{}, errors.New("Link not found or already deleted")
		}
		return []models.Links{}, errors.New("Failed to fetch link: " + err.Error())
	}

	return link, nil
}

func (u LinkRepository) DeleteLinkById(id int) (models.Links, error) {
	sql := "UPDATE links SET deleted_at = $1 WHERE id = $2 AND deleted_at IS NULL RETURNING id, user_id, original_url, short_url, created_at, updated_at, deleted_at"

	rows, err := u.db.Query(context.Background(), sql, time.Now(), id)
	if err != nil {
		return models.Links{}, errors.New("Failed to delete link! : " + err.Error())
	}

	deletedLink, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Links])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.Links{}, errors.New("Link not found or already deleted")
		}

		return models.Links{}, errors.New("Link deleted but returning error! : " + err.Error())
	}

	return deletedLink, nil
}
