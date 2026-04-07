package services

import (
	"errors"
	"snowfoxinfinity/infinity-shortcut/internal/dto"
	"snowfoxinfinity/infinity-shortcut/internal/lib"
	"snowfoxinfinity/infinity-shortcut/internal/models"
	"snowfoxinfinity/infinity-shortcut/internal/repository"
	"strings"

	"github.com/matthewhartstonge/argon2"
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

	argon := argon2.DefaultConfig()

	hashedPassword, err := argon.HashEncoded([]byte(newUser.Password))
	if err != nil {
		return dto.RegisterResponseDTO{}, err
	}

	modeledNewUser := models.User{
		FirstName:    newUser.FirstName,
		LastName:     newUser.LastName,
		Email:        newUser.Email,
		PasswordHash: string(hashedPassword),
	}

	_, err = u.userRepo.GetUserByEmail(modeledNewUser.Email)
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

func (a AuthService) Login(req dto.LoginRequestDTO) (dto.LoginResponseDTO, error) {
	if !strings.Contains(req.Email, "@") {
		return dto.LoginResponseDTO{}, errors.New("Failed to login! Invalid email format !")
	}

	user, err := a.userRepo.GetUserByEmail(req.Email)
	if err != nil {
		return dto.LoginResponseDTO{}, errors.New("Failed to login! Invalid email or password !")
	}

	matched, err := argon2.VerifyEncoded([]byte(req.Password), []byte(user.PasswordHash))
	if err != nil {
		return dto.LoginResponseDTO{}, errors.New("Failed to login! Invalid email or password !")
	}

	if(!matched){
		return dto.LoginResponseDTO{}, errors.New("Failed to login! Invalid email or password !")
	}

	token, err := lib.GenerateToken(user.Id)
	if err != nil {
		return dto.LoginResponseDTO{}, errors.New("Failed to generate token : " + err.Error())
	}

	return dto.LoginResponseDTO{
		Token: token,
	}, nil
}
