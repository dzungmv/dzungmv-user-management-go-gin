package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorCode string

const (
	ErrCodeBadRequest ErrorCode = "BAD_REQUEST"
	ErrCodeNotFound   ErrorCode = "NOT_FOUND"
	ErrCodeConflict   ErrorCode = "CONFLICT"
	ErrCodeInternal   ErrorCode = "INTERNAL_SERVER_ERROR"
)

type AppError struct {
	Message string
	Code    string
	Err     error
}

func (ae *AppError) Error() string {
	return ""
}

func NewError(message string, code string) error {
	return &AppError{
		Message: message,
		Code:    code,
	}
}

func WrapError(err error, message string, code string) error {
	return &AppError{
		Err:     err,
		Message: message,
		Code:    code,
	}
}

func ResponseSuccess(ctx *gin.Context, status int, data any) {
	ctx.JSON(status, gin.H{
		"status": "success",
		"data":   data,
	})
}

func ResponseError(ctx *gin.Context, err error) {
	if appErr, ok := err.(*AppError); ok {

		status := httpStatusFromCode(ErrorCode(appErr.Code))

		response := gin.H{
			"error": appErr.Message,
			"code":  appErr.Code,
		}

		if appErr.Err != nil {
			response["detail"] = appErr.Err.Error()
		}

		ctx.JSON(status, response)
		return
	}

	ctx.JSON(http.StatusInternalServerError, gin.H{
		"error": err.Error(),
		"code":  http.StatusInternalServerError,
	})

}

func httpStatusFromCode(code ErrorCode) int {
	switch code {
	case ErrCodeBadRequest:
		return http.StatusBadRequest
	case ErrCodeNotFound:
		return http.StatusNotFound
	case ErrCodeConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
