package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Todo struct {
	ID          string         `json:"id" gorm:"primaryKey" example:"550e8400-e29b-41d4-a716-446655440000"`
	Title       string         `json:"title" gorm:"not null" example:"Buy groceries" validate:"required,min=1,max=200"`
	Description string         `json:"description" example:"Milk, eggs, and bread" validate:"max=1000"`
	Completed   bool           `json:"completed" gorm:"default:false" example:"false"`
	CreatedAt   time.Time      `json:"created_at" example:"2026-04-25T14:30:00Z"`
	UpdatedAt   time.Time      `json:"updated_at" example:"2026-04-25T14:30:00Z"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index" swaggerignore:"true"`
}

func (t *Todo) BeforeCreate(tx *gorm.DB) error {
	if t.ID == "" {
		t.ID = uuid.New().String()
	}
	return nil
}

type CreateTodoRequest struct {
	Title       string `json:"title" example:"Learn Fiber v3" validate:"required,min=1,max=200"`
	Description string `json:"description" example:"Build a CRUD app with Swagger" validate:"max=1000"`
}

type UpdateTodoRequest struct {
	Title       *string `json:"title,omitempty" example:"Updated title" validate:"omitempty,min=1,max=200"`
	Description *string `json:"description,omitempty" example:"Updated description" validate:"omitempty,max=1000"`
	Completed   *bool   `json:"completed,omitempty" example:"true"`
}
