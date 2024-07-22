package util

import (
	"bae-backend/internal/core/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

// response represents a response body format
type response struct {
	Success bool   `json:"success" example:"true"`
	Message string `json:"message" example:"Success"`
	Data    any    `json:"data,omitempty"`
}

// newResponse is a helper function to create a response body
func newResponse(success bool, message string, data any) response {
	return response{
		Success: success,
		Message: message,
		Data:    data,
	}
}

// errorResponse represents an error response body format
type errorResponse struct {
	Success  bool     `json:"success" example:"false"`
	Messages []string `json:"messages" example:"Error message 1, Error message 2"`
}

// newErrorResponse is a helper function to create an error response body
func newErrorResponse(errMsgs []string) errorResponse {
	return errorResponse{
		Success:  false,
		Messages: errMsgs,
	}
}

// HandleSuccess sends a success response with the specified status code and optional data
func HandleSuccess(ctx *gin.Context, data any) {
	rsp := newResponse(true, "Success", data)
	ctx.JSON(http.StatusOK, rsp)
}

// HandleError send error response
func HandleError(ctx *gin.Context, err error) {
	statusCode, ok := domain.ErrorStatusMap[err]
	if !ok {
		statusCode = http.StatusInternalServerError
	}

	errMsg := err.Error()
	errRsp := newErrorResponse([]string{errMsg})
	ctx.JSON(statusCode, errRsp)
}
