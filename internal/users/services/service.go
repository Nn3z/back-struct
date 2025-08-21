package services

import (
	dtos "bazar/internal/users/DTOs"
	"bazar/internal/users/repositories"
	"errors"
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
