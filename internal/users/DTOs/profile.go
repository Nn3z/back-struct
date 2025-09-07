package dtos

type UserProfileDTO struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	MatSummary string `json:"mat_summary"`
	PatSummary string `json:"pat_summary"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	RoleID     int    `json:"role_id"`
	Phone      string `json:"phone"`
	Image      string `json:"image"`
}
