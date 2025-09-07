package services

import (
	dtos "bazar/internal/users/DTOs"
	"bazar/internal/users/repositories"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user *dtos.UserSignUpDTO) error {
	if user.Name == "" || user.MatSummary == "" || user.PatSummary == "" || user.Email == "" || user.Password == "" {
		return errors.New("missing required fields")
	}
	return s.repo.Create(user)
}

func (s *UserService) Login(req dtos.UserLoginDTO) (*dtos.LoginResponse, error) {
	user, err := s.repo.FindByEmail(req.Email)

	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)) != nil {
		return nil, errors.New("invalid credentials")
	}

	claims := jwt.MapClaims{
		"id":     user.ID,
		"email":  user.Email,
		"RoleID": user.RoleID,
		"exp":    time.Now().Add(time.Hour * 168).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("JWT_SECRET")
	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return nil, err
	}
	return &dtos.LoginResponse{Token: signedToken}, nil
}

func (s *UserService) GetProfile(userID string) (*dtos.UserProfileDTO, error) {
	user, err := s.repo.FindByID(userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
