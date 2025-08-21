package services

import (
	"bazar/internal/role/model"
	"bazar/internal/role/repositories"
	"database/sql"
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

func (s *RoleService) UpdateServiceRole(role *model.Role) error {
	if role.ID == nil {
		return errors.New("ID are required")
	}
	return s.repo.Update(role)
}

func (s *RoleService) DeleteServiceRole(role *model.Role) error {
	if role.ID == nil {
		return errors.New("ID are required")
	}
	var existsID string
	sqlStatement := `SELECT id FROM role WHERE id = $1`
	err := s.repo.DB.QueryRow(sqlStatement, role.ID).Scan(&existsID)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("role not found")
		}
		return err
	}
	return s.repo.Delete(role)
}

func (s *RoleService) GetServiceRole(role *model.Role) error {
	sqlStatement := `SELECT * FROM role`
	err := s.repo.DB.QueryRow(sqlStatement).Scan(&role.ID)
	if err == nil {
		if err == sql.ErrNoRows {
			return errors.New("role not found")
		}
		return err
	}
	return s.repo.Delete(role)
}
