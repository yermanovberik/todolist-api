package dto

type TaskRequestDTO struct {
	Title    string `json:"title"`
	ActiveAt string `json:"activeAt"`
}
