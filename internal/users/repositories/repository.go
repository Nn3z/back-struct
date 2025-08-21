package repositories

import (
	dtos "bazar/internal/users/DTOs"
	"bazar/pkg/utils"
	"database/sql"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) Create(user *dtos.UserSignUpDTO) error {
	data := map[string]interface{}{
		"id":          user.ID,
		"name":        user.Name,
		"mat_summary": user.MatSummary,
		"pat_summary": user.PatSummary,
		"username":    user.Username,
		"email":       user.Email,
		"password":    user.Password,
		"phone":       user.Phone,
		"image":       user.Image,
		"role_id":     user.RoleID,
	}
	query, values := utils.ParseInsertArray("users", data)
	_, err := r.DB.Exec(query, values...)
	return err
}

//pendiente de hacer   "login"
