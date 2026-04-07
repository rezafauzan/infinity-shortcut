package services

import (
	"errors"
	"snowfoxinfinity/infinity-shortcut/internal/dto"
	"snowfoxinfinity/infinity-shortcut/internal/lib"
	"snowfoxinfinity/infinity-shortcut/internal/models"
	"snowfoxinfinity/infinity-shortcut/internal/repository"
)

type LinkService struct {
	linkRepo *repository.LinkRepository
}

func NewLinkService(linkRepo *repository.LinkRepository) *LinkService {
	return &LinkService{
		linkRepo: linkRepo,
	}
}

func (l LinkService) CreateNewLink(newLink dto.CreateNewLinkDTO, userId int) (dto.CreateNewLinkResponseDTO, error) {
	if userId < 0 {
		return dto.CreateNewLinkResponseDTO{}, errors.New("Session invalid please relogin!")
	}
	
	slug := lib.Shuffle()

	modeledNewLink := models.Links{
		UserId:      userId,
		OriginalUrl: newLink.OriginalUrl,
		ShortUrl:    "localhost:8888/api/links/" + slug,
	}
	
	createdNewLink, err := l.linkRepo.CreateNewLink(modeledNewLink)
	if err != nil {
		return dto.CreateNewLinkResponseDTO{}, err
	}

	response := dto.CreateNewLinkResponseDTO{
		Id:          createdNewLink.Id,
		UserId:      createdNewLink.UserId,
		OriginalUrl: createdNewLink.OriginalUrl,
		ShortUrl:    createdNewLink.ShortUrl,
		CreatedAt:   createdNewLink.CreatedAt,
		UpdatedAt:   createdNewLink.UpdatedAt,
	}

	return response, nil
}
