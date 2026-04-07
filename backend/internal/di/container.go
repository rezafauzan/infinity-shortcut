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
	AuthHandler *handlers.AuthHandler
	LinkHandler *handlers.LinkHandler
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
	userService := services.NewAuthService(userRepo)
	c.AuthHandler = handlers.NewAuthHandler(userService)
	
	linkRepo, err := repository.NewLinkRepository(c.db)
	if err != nil {
		return err
	}
	linkService := services.NewLinkService(linkRepo)
	c.LinkHandler = handlers.NewLinkHandler(linkService)

	return nil
}