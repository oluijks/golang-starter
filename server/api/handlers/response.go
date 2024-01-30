package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type APIResponse struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

func ResponseCreated(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "201 Created",
	})
}

func ResponseBadRequest(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusBadRequest, APIResponse{
		Status:  http.StatusBadRequest,
		Message: "400 Bad Request",
		Data:    map[string]interface{}{"error": err.Error()},
	})
}

func ResponseUnAuthorized(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusUnauthorized, APIResponse{
		Status:  http.StatusUnauthorized,
		Message: "401 Unauthorized",
		Data:    map[string]interface{}{"error": err.Error()},
	})
}

func ResponseNotFound(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusNotFound, APIResponse{
		Status:  http.StatusNotFound,
		Message: "404 Not Found",
		Data:    map[string]interface{}{"error": err.Error()},
	})
}

func ResponseServerError(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusInternalServerError, APIResponse{
		Status:  http.StatusInternalServerError,
		Message: "500 Internal Server Error",
		Data:    map[string]interface{}{"error": err.Error()},
	})
}
