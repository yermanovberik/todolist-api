package domain

import "time"

type Task struct {
	ID       int       `json:"id"`
	Tittle   string    `json:"tittle"`
	ActiveAt time.Time `json:"activeAt"`
	Done     bool      `json:"done"`
}
