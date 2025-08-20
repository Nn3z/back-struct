package repositories

import (
	"bazar/internal/role/model"
	"bazar/pkg/utils"
	"database/sql"
)

type RoleRepository struct {
	DB *sql.DB
}

func NewRoleRepository(db *sql.DB) *RoleRepository {
	return &RoleRepository{DB: db}
}

func (r *RoleRepository) Create(role *model.Role) error {
	data := map[string]interface{}{
		"id":          role.ID,
		"title":       role.Title,
		"description": role.Description,
	}
	query, values := utils.ParseInsertArray("role", data)
	_, err := r.DB.Exec(query, values...)
	return err
}
