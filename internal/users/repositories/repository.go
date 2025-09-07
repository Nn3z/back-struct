package repositories

import (
	dtos "bazar/internal/users/DTOs"
	"bazar/internal/users/model"
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

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User
	query := `SELECT id, name, email, password, role_id FROM users WHERE LOWER(email) = LOWER($1) LIMIT 1`
	err := r.DB.QueryRow(query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.RoleID)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindByID(id string) (*dtos.UserProfileDTO, error) {
	var user dtos.UserProfileDTO
	query := `SELECT id, name, mat_summary, pat_summary, email, username, role_id, phone, image FROM users WHERE id = $1 LIMIT 1`
	err := r.DB.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.MatSummary, &user.PatSummary, &user.Email, &user.Username, &user.RoleID, &user.Phone, &user.Image)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
