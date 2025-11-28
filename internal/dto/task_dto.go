package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateTaskRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description,omitempty"`
}

type UpdateTaskRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description,omitempty" validate:"omitempty"`
	Status      int    `json:"status"`
}
type TaskResponse struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description,omitempty"`
	Status      int       `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
