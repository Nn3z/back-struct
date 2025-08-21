package dtos

type UserSignUpDTO struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	MatSummary string `json:"mat_summary"`
	PatSummary string `json:"pat_summary"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Phone      string `json:"phone"`
	Image      string `json:"image"`
	RoleID     string `json:"role_id"`
}
