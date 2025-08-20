package services

import (
	"bazar/internal/role/model"
	"bazar/internal/role/repositories"
	"errors"
)

type RoleService struct {
	repo *repositories.RoleRepository
}

func NewRoleService(repo *repositories.RoleRepository) *RoleService {
	return &RoleService{repo: repo}
}

func (s *RoleService) CreateRole(role *model.Role) error {
	if role.ID == nil || role.Title == "" || role.Description == "" {
		return errors.New("ID, TITLE AND DESCRIPTION ARE REQUIRED")
	}
	return s.repo.Create(role)
}
