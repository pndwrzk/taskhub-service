package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pndwrzk/taskhub-service/internal/constants"
	errConst "github.com/pndwrzk/taskhub-service/internal/constants/error"
)

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ParamHTTPResp struct {
	Code     int
	Error    error
	Messsage *string
	Gin      *gin.Context
	Data     interface{}
}

func HttpResponse(param ParamHTTPResp) {
	if param.Error == nil {
		param.Gin.JSON(param.Code, Response{
			Status:  constants.AppSuccess,
			Message: http.StatusText(http.StatusOK),
			Data:    param.Data,
		})
		return
	}

	message := getErrorMessage(param)

	param.Gin.JSON(param.Code, Response{
		Status:  constants.AppError,
		Message: message,
		Data:    param.Data,
	})
}

func getErrorMessage(param ParamHTTPResp) string {
	if param.Messsage != nil {
		return *param.Messsage
	}
	if param.Error != nil && errConst.ErrMapping(param.Error) {
		return param.Error.Error()
	}
	return errConst.ErrInternalServerError.Error()
}
