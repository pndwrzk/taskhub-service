// internal/usecase/task_usecase.go
package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/pndwrzk/taskhub-service/internal/constants"
	"github.com/pndwrzk/taskhub-service/internal/dto"
	"github.com/pndwrzk/taskhub-service/internal/model"
	"github.com/pndwrzk/taskhub-service/internal/repository"
)

type TaskUsecase interface {
	CreateTask(ctx context.Context, req *dto.CreateTaskRequest, userID *uuid.UUID) (*dto.TaskResponse, error)
	GetTasksByUser(ctx context.Context, userID *uuid.UUID) ([]*dto.TaskResponse, error)
	UpdateTask(ctx context.Context, req *dto.UpdateTaskRequest, id *uuid.UUID) (*dto.TaskResponse, error)
	DeleteTask(ctx context.Context, id *uuid.UUID) error
}

type taskUsecase struct {
	taskRepo repository.TaskRepository
}

func NewTaskUsecase(taskRepo repository.TaskRepository) TaskUsecase {
	return &taskUsecase{taskRepo: taskRepo}
}

func (u *taskUsecase) CreateTask(ctx context.Context, req *dto.CreateTaskRequest, userID *uuid.UUID) (*dto.TaskResponse, error) {
	task := &model.Task{
		ID:          uuid.New(),
		Title:       req.Title,
		Description: req.Description,
		Status:      constants.TaskTodo,
		UserID:      *userID,
	}

	if err := u.taskRepo.Create(ctx, task); err != nil {
		return nil, err
	}

	return &dto.TaskResponse{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}, nil
}

func (u *taskUsecase) GetTasksByUser(ctx context.Context, userID *uuid.UUID) ([]*dto.TaskResponse, error) {
	tasks, err := u.taskRepo.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	var res []*dto.TaskResponse
	for _, t := range tasks {
		res = append(res, &dto.TaskResponse{
			ID:          t.ID,
			Title:       t.Title,
			Description: t.Description,
			Status:      t.Status,
			CreatedAt:   t.CreatedAt,
			UpdatedAt:   t.UpdatedAt,
		})
	}

	return res, nil
}

func (u *taskUsecase) UpdateTask(ctx context.Context, req *dto.UpdateTaskRequest, id *uuid.UUID) (*dto.TaskResponse, error) {
	tasks, err := u.taskRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	task := &model.Task{
		ID:          *id,
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
		UserID:      tasks.UserID,
		UpdatedAt:   time.Now(),
	}

	if err := u.taskRepo.Update(ctx, task); err != nil {
		return nil, err
	}

	return &dto.TaskResponse{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		UpdatedAt:   task.UpdatedAt,
	}, nil
}

func (u *taskUsecase) DeleteTask(ctx context.Context, id *uuid.UUID) error {
	if err := u.taskRepo.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}
