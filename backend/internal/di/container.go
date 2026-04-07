package di

import (
	"snowfoxinfinity/infinity-shortcut/internal/handlers"
	"snowfoxinfinity/infinity-shortcut/internal/lib"
	"snowfoxinfinity/infinity-shortcut/internal/repository"
	"snowfoxinfinity/infinity-shortcut/internal/services"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Container struct {
	db          *pgxpool.Pool
	UserHandler *handlers.UserHandler
}

func NewContainer() (*Container, error) {
	db, err := lib.DatabaseConnect()
	if err != nil {
		return nil, err
	}

	container := &Container{
		db: db,
	}

	container.initDependencies()
	
	return container, nil
}

func (c *Container) initDependencies() error {
	userRepo, err := repository.NewUserRepository(c.db)
	if err != nil {
		return err
	}
	userService := services.NewUserService(userRepo)
	c.UserHandler = handlers.NewUserHandler(userService)

	return nil
}