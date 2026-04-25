package storage

import (
	"errors"
	"todo-api/database"
	models "todo-api/model"

	"gorm.io/gorm"
)

var (
	ErrTodoNotFound = errors.New("todo not found")
)

type SQLiteStore struct{}

func NewSQLiteStore() *SQLiteStore {
	return &SQLiteStore{}
}

func (s *SQLiteStore) Create(todo *models.Todo) (*models.Todo, error) {
	if err := database.DB.Create(todo).Error; err != nil {
		return nil, err
	}
	return todo, nil
}

func (s *SQLiteStore) GetAll() ([]models.Todo, error) {
	var todos []models.Todo
	if err := database.DB.Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}

func (s *SQLiteStore) GetByID(id string) (*models.Todo, error) {
	var todo models.Todo
	if err := database.DB.First(&todo, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrTodoNotFound
		}
		return nil, err
	}
	return &todo, nil
}

func (s *SQLiteStore) Update(id string, req *models.UpdateTodoRequest) (*models.Todo, error) {
	todo, err := s.GetByID(id)
	if err != nil {
		return nil, err
	}

	updates := map[string]interface{}{}
	if req.Title != nil {
		updates["title"] = *req.Title
	}
	if req.Description != nil {
		updates["description"] = *req.Description
	}
	if req.Completed != nil {
		updates["completed"] = *req.Completed
	}

	if len(updates) > 0 {
		if err := database.DB.Model(todo).Updates(updates).Error; err != nil {
			return nil, err
		}
	}

	return todo, nil
}

func (s *SQLiteStore) Delete(id string) error {
	result := database.DB.Delete(&models.Todo{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrTodoNotFound
	}
	return nil
}
