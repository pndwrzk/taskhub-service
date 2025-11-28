// internal/delivery/http/handler/task_handler.go
package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	errWap "github.com/pndwrzk/taskhub-service/internal/common/error"
	"github.com/pndwrzk/taskhub-service/internal/common/response"
	"github.com/pndwrzk/taskhub-service/internal/dto"
	"github.com/pndwrzk/taskhub-service/internal/usecase"
)

type TaskHandler struct {
	taskUsecase usecase.TaskUsecase
}

func NewTaskHandler(taskUsecase usecase.TaskUsecase) *TaskHandler {
	return &TaskHandler{taskUsecase: taskUsecase}
}

func (h *TaskHandler) CreateTask(ctx *gin.Context) {
	userIDStr := ctx.GetString("user_id")
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code:  http.StatusUnauthorized,
			Gin:   ctx,
			Error: err,
		})
		return
	}

	var req dto.CreateTaskRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code:  http.StatusBadRequest,
			Gin:   ctx,
			Error: err,
		})
		return
	}

	validate := validator.New()
	err = validate.Struct(req)

	if err != nil {
		errMessage := http.StatusText(http.StatusUnprocessableEntity)
		errResponse := errWap.WrapError(err)
		errData := errWap.ErrValidationResponse(err)
		response.HttpResponse(response.ParamHTTPResp{
			Code:     http.StatusUnprocessableEntity,
			Gin:      ctx,
			Error:    errResponse,
			Data:     errData,
			Messsage: &errMessage,
		})

		return
	}

	task, err := h.taskUsecase.CreateTask(ctx, &req, &userID)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code:  http.StatusInternalServerError,
			Gin:   ctx,
			Error: err,
		})
		return
	}

	response.HttpResponse(response.ParamHTTPResp{
		Code: http.StatusCreated,
		Gin:  ctx,
		Data: task,
	})
}

func (h *TaskHandler) GetTasksByUser(ctx *gin.Context) {
	userIDStr := ctx.GetString("user_id")
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code:  http.StatusUnauthorized,
			Gin:   ctx,
			Error: err,
		})
		return
	}

	tasks, err := h.taskUsecase.GetTasksByUser(ctx, &userID)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code:  http.StatusInternalServerError,
			Gin:   ctx,
			Error: err,
		})
		return
	}

	response.HttpResponse(response.ParamHTTPResp{
		Code: http.StatusOK,
		Gin:  ctx,
		Data: tasks,
	})
}

func (h *TaskHandler) UpdateTask(ctx *gin.Context) {

	taskIDStr := ctx.Param("id")
	taskID, err := uuid.Parse(taskIDStr)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code:  http.StatusBadRequest,
			Gin:   ctx,
			Error: err,
		})
		return
	}

	var req dto.UpdateTaskRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code:  http.StatusBadRequest,
			Gin:   ctx,
			Error: err,
		})
		return
	}

	validate := validator.New()
	err = validate.Struct(req)

	if err != nil {
		errMessage := http.StatusText(http.StatusUnprocessableEntity)
		errResponse := errWap.WrapError(err)
		errData := errWap.ErrValidationResponse(err)
		response.HttpResponse(response.ParamHTTPResp{
			Code:     http.StatusUnprocessableEntity,
			Gin:      ctx,
			Error:    errResponse,
			Data:     errData,
			Messsage: &errMessage,
		})

		return
	}

	task, err := h.taskUsecase.UpdateTask(ctx, &req, &taskID)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code:  http.StatusInternalServerError,
			Gin:   ctx,
			Error: err,
		})
		return
	}

	response.HttpResponse(response.ParamHTTPResp{
		Code: http.StatusOK,
		Gin:  ctx,
		Data: task,
	})
}

func (h *TaskHandler) DeleteTask(ctx *gin.Context) {
	taskIDStr := ctx.Param("id")
	taskID, err := uuid.Parse(taskIDStr)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code:  http.StatusBadRequest,
			Gin:   ctx,
			Error: err,
		})
		return
	}

	if err := h.taskUsecase.DeleteTask(ctx, &taskID); err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code:  http.StatusInternalServerError,
			Gin:   ctx,
			Error: err,
		})
		return
	}

	response.HttpResponse(response.ParamHTTPResp{
		Code: http.StatusNoContent,
		Gin:  ctx,
	})
}
