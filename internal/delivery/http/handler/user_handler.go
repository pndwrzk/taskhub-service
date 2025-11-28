package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	errWap "github.com/pndwrzk/taskhub-service/internal/common/error"
	"github.com/pndwrzk/taskhub-service/internal/common/response"
	"github.com/pndwrzk/taskhub-service/internal/dto"
	"github.com/pndwrzk/taskhub-service/internal/usecase"
)

type UserHandler struct {
	usecase usecase.UserUsecase
}

func NewUserHandler(usecase usecase.UserUsecase) *UserHandler {
	return &UserHandler{usecase: usecase}
}

func (h *UserHandler) Register(ctx *gin.Context) {
	var request dto.RegisterRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code:  http.StatusBadRequest,
			Gin:   ctx,
			Error: err,
		})
		return
	}

	validate := validator.New()
	err := validate.Struct(request)

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

	res, err := h.usecase.Register(ctx, &request)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code:  http.StatusBadRequest,
			Gin:   ctx,
			Error: err,
		})
		return
	}

	response.HttpResponse(response.ParamHTTPResp{
		Code: http.StatusCreated,
		Data: res,
		Gin:  ctx,
	})

}

func (h *UserHandler) Login(ctx *gin.Context) {
	var request dto.LoginRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code:  http.StatusBadRequest,
			Gin:   ctx,
			Error: err,
		})
		return
	}

	validate := validator.New()
	err := validate.Struct(request)

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

	res, err := h.usecase.Login(ctx, &request)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code:  http.StatusBadRequest,
			Gin:   ctx,
			Error: err,
		})
		return
	}

	response.HttpResponse(response.ParamHTTPResp{
		Code: http.StatusOK,
		Data: res,
		Gin:  ctx,
	})

}

func (h *UserHandler) RefreshToken(ctx *gin.Context) {
	var request dto.RefreshTokenRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code:  http.StatusBadRequest,
			Gin:   ctx,
			Error: err,
		})
		return
	}

	validate := validator.New()
	err := validate.Struct(request)

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

	res, err := h.usecase.RefreshToken(ctx, &request)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code:  http.StatusBadRequest,
			Gin:   ctx,
			Error: err,
		})
		return
	}

	response.HttpResponse(response.ParamHTTPResp{
		Code: http.StatusOK,
		Data: res,
		Gin:  ctx,
	})

}
