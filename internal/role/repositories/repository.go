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

func (r *RoleRepository) Update(role *model.Role) error {
	data := map[string]interface{}{
		"id":          role.ID,
		"title":       role.Title,
		"description": role.Description,
	}
	condition := map[string]interface{}{
		"id": role.ID,
	}

	query, values := utils.ParseUpdateArray("role", data, condition)
	_, err := r.DB.Exec(query, values...)
	return err
}

func (r *RoleRepository) Delete(role *model.Role) error {
	query := `DELETE FROM role WHERE id = $1`
	_, err := r.DB.Exec(query, role.ID)
	return err
}

func (r *RoleRepository) Get(role *model.Role) error {
	query := `SELECT * FROM role`
	_, err := r.DB.Query(query)
	return err
}
