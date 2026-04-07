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

func (l LinkService) GetLinkById(linkId int) (dto.GetLinkResponseDTO, error) {
	if linkId < 0 {
		return dto.GetLinkResponseDTO{}, errors.New("Link data invalid !")
	}

	link, err := l.linkRepo.GetLinkById(linkId)
	if err != nil {
		return dto.GetLinkResponseDTO{}, err
	}

	response := dto.GetLinkResponseDTO{
		Id:          link.Id,
		UserId:      link.UserId,
		OriginalUrl: link.OriginalUrl,
		ShortUrl:    link.ShortUrl,
		CreatedAt:   link.CreatedAt,
		UpdatedAt:   link.UpdatedAt,
	}

	return response, nil
}

func (l LinkService) GetAllLinksByUserId(userId int) ([]dto.GetLinkResponseDTO, error) {
	if userId < 0 {
		return []dto.GetLinkResponseDTO{}, errors.New("Session invalid please relogin!")
	}

	links, err := l.linkRepo.GetAllLinksByUserId(userId)
	if err != nil {
		return []dto.GetLinkResponseDTO{}, err
	}

	var response []dto.GetLinkResponseDTO

	for _, link := range links {
		response = append(response, dto.GetLinkResponseDTO{
			Id:          link.Id,
			UserId:      link.UserId,
			OriginalUrl: link.OriginalUrl,
			ShortUrl:    link.ShortUrl,
			CreatedAt:   link.CreatedAt,
			UpdatedAt:   link.UpdatedAt,
		})
	}

	return response, nil

}

func (l LinkService) DeleteLinkById(id int, userId int) (dto.DeleteLinkResponseDTO, error) {
	if userId < 0 {
		return dto.DeleteLinkResponseDTO{}, errors.New("Session invalid please relogin!")
	}

	link, err := l.linkRepo.GetLinkById(id)
	if err != nil {
		return dto.DeleteLinkResponseDTO{}, errors.New("Link not found!")
	}

	if link.UserId != userId {
		return dto.DeleteLinkResponseDTO{}, errors.New("Forbidden: you don't have access to this link!")
	}

	deletedLink, err := l.linkRepo.DeleteLinkById(id)
	if err != nil {
		return dto.DeleteLinkResponseDTO{}, err
	}

	response := dto.DeleteLinkResponseDTO{
		Id:          deletedLink.Id,
		UserId:      deletedLink.UserId,
		OriginalUrl: deletedLink.OriginalUrl,
		ShortUrl:    deletedLink.ShortUrl,
		DeletedAt:   *deletedLink.DeletedAt,
	}

	return response, nil
}