package services

import (
	"errors"
	"snowfoxinfinity/infinity-shortcut/internal/dto"
	"snowfoxinfinity/infinity-shortcut/internal/models"
	"snowfoxinfinity/infinity-shortcut/internal/repository"
	"strings"
)

type AuthService struct {
	userRepo *repository.UserRepository
}

func NewAuthService(userRepo *repository.UserRepository) *AuthService {
	return &AuthService{
		userRepo: userRepo,
	}
}

func (u AuthService) Register(newUser dto.RegisterDTO) (dto.RegisterResponseDTO, error) {
	if len(newUser.FirstName) < 4 {
		return dto.RegisterResponseDTO{}, errors.New("First name length minimum is 4 characters !")
	}
	if len(newUser.LastName) < 4 {
		return dto.RegisterResponseDTO{}, errors.New("Last name length minimum is 4 characters !")
	}
	if !strings.Contains(newUser.Email, "@") {
		return dto.RegisterResponseDTO{}, errors.New("Invalid email format !")
	}
	if len(newUser.Password) < 8 {
		return dto.RegisterResponseDTO{}, errors.New("Password too weak minimum length is 8 characters !")
	}
	if newUser.PasswordConfirm != newUser.Password {
		return dto.RegisterResponseDTO{}, errors.New("Password confirmation missmatch !")
	}

	modeledNewUser := models.User{
		FirstName:    newUser.FirstName,
		LastName:     newUser.LastName,
		Email:        newUser.Email,
		PasswordHash: newUser.Password,
	}

	_, err := u.userRepo.GetUserByEmail(modeledNewUser.Email)
	if err == nil {
		return dto.RegisterResponseDTO{}, errors.New("Email allready used!")
	}

	registeredUser, err := u.userRepo.CreateNewUser(modeledNewUser)
	if err != nil {
		return dto.RegisterResponseDTO{}, err
	}

	response := dto.RegisterResponseDTO{
		FirstName: registeredUser.FirstName,
		LastName:  registeredUser.LastName,
		Email:     registeredUser.Email,
		CreatedAt: registeredUser.CreatedAt,
		UpdatedAt: registeredUser.UpdatedAt,
	}

	return response, nil
}
