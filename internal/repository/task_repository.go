// internal/repository/task_repository.go
package repository

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	errConst "github.com/pndwrzk/taskhub-service/internal/constants/error"
	"github.com/pndwrzk/taskhub-service/internal/model"
	"gorm.io/gorm"
)

type TaskRepository interface {
	Create(ctx context.Context, task *model.Task) error
	GetByUserID(ctx context.Context, userID *uuid.UUID) ([]*model.Task, error)
	GetByID(ctx context.Context, id *uuid.UUID) (*model.Task, error)
	Update(ctx context.Context, task *model.Task) error
	Delete(ctx context.Context, id *uuid.UUID) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) Create(ctx context.Context, task *model.Task) error {
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()
	err := r.db.WithContext(ctx).Create(task).Error
	if err != nil {
		return errConst.ErrSqlError
	}
	return nil
}

func (r *taskRepository) GetByUserID(ctx context.Context, userID *uuid.UUID) ([]*model.Task, error) {
	var tasks []*model.Task
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&tasks).Error
	if err != nil {
		return nil, errConst.ErrSqlError
	}
	return tasks, nil
}

func (r *taskRepository) Update(ctx context.Context, task *model.Task) error {
	task.UpdatedAt = time.Now()
	err := r.db.WithContext(ctx).Save(task).Error
	if err != nil {
		return errConst.ErrSqlError
	}
	return nil
}

func (r *taskRepository) Delete(ctx context.Context, id *uuid.UUID) error {
	err := r.db.WithContext(ctx).Delete(&model.Task{}, id).Error
	if err != nil {
		return errConst.ErrSqlError
	}
	return nil
}

func (r *taskRepository) GetByID(ctx context.Context, id *uuid.UUID) (*model.Task, error) {
	var task model.Task
	err := r.db.WithContext(ctx).First(&task, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errConst.ErrTaskNotFound
		}
		return nil, errConst.ErrSqlError
	}
	return &task, nil
}
