package services

import (
	"errors"
	"snowfoxinfinity/infinity-shortcut/internal/dto"
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

func (l LinkService) CreateNewLink(newLink dto.CreateNewLinkDTO) (dto.CreateNewLinkResponseDTO, error) {
	if newLink.UserId < 0 {
		return dto.CreateNewLinkResponseDTO{}, errors.New("Session invalid please relogin!")
	}

	modeledNewLink := models.Links{
		UserId:      newLink.UserId,
		OriginalUrl: newLink.OriginalUrl,
		ShortUrl:    newLink.ShortUrl,
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
