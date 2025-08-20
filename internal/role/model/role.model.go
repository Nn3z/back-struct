package model

type Role struct {
	ID          *string `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
}
