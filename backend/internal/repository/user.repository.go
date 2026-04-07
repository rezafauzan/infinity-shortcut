package repository

import (
	"context"
	"errors"
	"snowfoxinfinity/infinity-shortcut/internal/models"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) (*UserRepository, error) {
	return &UserRepository{
		db: db,
	}, nil
}

func (u UserRepository) CreateNewUser(newUser models.User) (models.User, error) {
	sql := "INSERT INTO users (first_name, last_name, email, password_hash, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, first_name, last_name, email, password_hash, created_at, updated_at"
	rows, err := u.db.Query(context.Background(), sql, newUser.FirstName, newUser.LastName, newUser.Email, newUser.PasswordHash, time.Now(), time.Now())
	if err != nil {
		return models.User{}, errors.New("Failed to create new user! : " + err.Error())
	}

	registeredUser, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.User])
	if err != nil {
		return models.User{}, errors.New("User created but returning error! : " + err.Error())
	}

	return registeredUser, nil
}