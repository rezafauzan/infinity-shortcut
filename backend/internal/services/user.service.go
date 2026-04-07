package services

import (
	"errors"
	"snowfoxinfinity/infinity-shortcut/internal/dto"
	"snowfoxinfinity/infinity-shortcut/internal/models"
	"snowfoxinfinity/infinity-shortcut/internal/repository"
	"strings"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (u UserService) CreateNewUser(newUser dto.CreateUserDTO) (dto.CreateUserResponseDTO, error) {
	if len(newUser.FirstName) < 4 {
		return dto.CreateUserResponseDTO{}, errors.New("Failed to create user! : First name length minimum is 4 characters !")
	}
	if len(newUser.LastName) < 4 {
		return dto.CreateUserResponseDTO{}, errors.New("Failed to create user! : Last name length minimum is 4 characters !")
	}
	if !strings.Contains(newUser.Email, "@") {
		return dto.CreateUserResponseDTO{}, errors.New("Failed to create user! : Invalid email format !")
	}
	if len(newUser.Password) < 8 {
		return dto.CreateUserResponseDTO{}, errors.New("Failed to create user! : Password too weak minimum length is 8 characters !")
	}
	if newUser.PasswordConfirm != newUser.Password {
		return dto.CreateUserResponseDTO{}, errors.New("Failed to create user! : Password confirmation missmatch !")
	}
	
	modeledNewUser := models.User{
		FirstName: newUser.FirstName,
		LastName:  newUser.LastName,
		Email:     newUser.Email,
		Password:  newUser.Password,
	}
	
	registeredUser, err := u.userRepo.CreateNewUser(modeledNewUser)
	if err != nil {
		return dto.CreateUserResponseDTO{}, errors.New("Failed to create new user! : Email allready used !")
	}
	
	response := dto.CreateUserResponseDTO{
		FirstName: registeredUser.FirstName,
		LastName: registeredUser.LastName,
		Email: registeredUser.Email,
		CreatedAt: registeredUser.CreatedAt,
		UpdatedAt: registeredUser.UpdatedAt,
	}

	return response, nil
}
