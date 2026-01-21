package services

import (
	"errors"
	"strings"
	"github.com/wahyujatirestu/simple-procurement-system/models"
	"github.com/wahyujatirestu/simple-procurement-system/repositories"
	"github.com/wahyujatirestu/simple-procurement-system/dto"
	"github.com/wahyujatirestu/simple-procurement-system/security"
	"github.com/wahyujatirestu/simple-procurement-system/utils/services"
)

type AuthService interface {
	Register(req dto.RegisterRequest) (*dto.AuthResponse, error)
	Login(req dto.LoginRequest) (*dto.AuthResponse, error)
}

type authService struct {
	userRepo repositories.UserRepository
	jwtService services.JwtService
}	

func NewAuthService(userRepo repositories.UserRepository, jwtService services.JwtService) AuthService {
	return &authService{userRepo: userRepo, jwtService: jwtService}
}	

func (s *authService) Register(req dto.RegisterRequest) (*dto.AuthResponse, error) {

	req.Username = strings.TrimSpace(strings.ToLower(req.Username))
	if req.Username == "" {
		return nil, errors.New("username is required")
	}

	if strings.Contains(req.Username, " ") {
		return nil, errors.New("username must not contain spaces")
	}

	existing, err := s.userRepo.FindByUsername(req.Username)
	if err == nil && existing != nil {
		return nil, errors.New("user already exists")
	}

	if req.Password != req.ConfirmPassword {
		return nil, errors.New("passwords do not match")
	}

	if len(req.Password) < 8 {
		return nil, errors.New("password must be at least 8 characters")
	}


	hash, err := security.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := models.User{
		Username: req.Username,
		Password: hash,
		Role:     req.Role,
	}

	if err := s.userRepo.Create(&user); err != nil {
		return nil, err
	}

	return &dto.AuthResponse{
		ID:       user.ID,
		Username: user.Username,
		Role:     user.Role,
	}, nil
}

func (s *authService) Login(req dto.LoginRequest) (*dto.AuthResponse, error) {
	user, err := s.userRepo.FindByUsername(req.Username)
	if err != nil {
		return nil, errors.New("invalid username or password")
	}

	if err := 	security.VerifyPassword(user.Password, req.Password); err != nil {
		return nil, errors.New("invalid username or password")
	}

	token, err := s.jwtService.CreateToken(*user)
	if err != nil {
		return nil, err
	}

	return &dto.AuthResponse{
		ID:       user.ID,
		Username: user.Username,
		Role:     user.Role,
		Token:    token,
	}, nil
	
}

