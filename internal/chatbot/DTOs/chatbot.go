package dtos

type AskRequestDTO struct {
	Question string `json:"question"`
}

type AskResponseDTO struct {
	Answer string `json:"answer"`
}
