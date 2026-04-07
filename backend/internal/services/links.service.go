package services

import (
	"errors"
	"snowfoxinfinity/infinity-shortcut/internal/dto"
	"snowfoxinfinity/infinity-shortcut/internal/models"
	"snowfoxinfinity/infinity-shortcut/internal/repository"
	"strings"
)

type LinkService struct {
	linkRepo *repository.LinkRepository
}

func NewLinkService(linkRepo *repository.LinkRepository) *LinkService {
	return &LinkService{
		linkRepo: linkRepo,
	}
}

func (u AuthService) CreateNewLink(newLink dto.CreateNewLinkDTO) (dto.CreateNewLinkResponseDTO, error) {
	if newLink.UserId < 0 {
		return dto.CreateNewLinkResponseDTO{}, errors.New("First name length minimum is 4 characters !")
	}

	if !strings.Contains(newLink.ShortUrl, "@") {
		return dto.CreateNewLinkResponseDTO{}, errors.New("Invalid email format !")
	}

	modeledNewLink := models.Links{
		UserId:      newLink.UserId,
		OriginalUrl: newLink.OriginalUrl,
		ShortUrl:    newLink.ShortUrl,
	}

	createNewLinkedUser, err := u.userRepo.CreateNewLink(modeledNewLink)
	if err != nil {
		return dto.CreateNewLinkResponseDTO{}, err
	}

	response := dto.CreateNewLinkResponseDTO{
		Id:          createNewLinkedUser.Id,
		UserId:      createNewLinkedUser.UserId,
		OriginalUrl: createNewLinkedUser.OriginalUrl,
		ShortUrl:    createNewLinkedUser.ShortUrl,
		CreatedAt:   createNewLinkedUser.CreatedAt,
		UpdatedAt:   createNewLinkedUser.UpdatedAt,
	}

	return response, nil
}
