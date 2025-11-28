package model

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	UserID      uuid.UUID `json:"user_id" gorm:"type:uuid;not null;index"`
	Title       string    `json:"title" gorm:"type:varchar(100);not null"`
	Description string    `json:"description" gorm:"type:text"`
	Status      int       `json:"status" gorm:"type:smallint;default:0;not null"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	User        User      `json:"user,omitempty" gorm:"foreignKey:UserID"`
}
