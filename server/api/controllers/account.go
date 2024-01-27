package controllers

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oluijks/golang-starter/server/api/models"
	"github.com/oluijks/golang-starter/server/internal/storage"
	"github.com/oluijks/golang-starter/server/internal/utils"
)

type AccountHandler struct {
	store storage.Store
}

func NewAccountHandlers(store storage.Store) *AccountHandler {
	return &AccountHandler{store: store}
}

type accountIDPayload struct {
	ID string `uri:"id" binding:"required,min=1"`
}

type createAccountPayload struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func (accountHandler *AccountHandler) CreateAccount(ctx *gin.Context) {
	var payload createAccountPayload
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ResponseBadRequest(ctx, err)
		return
	}

	hashedPassword, err := utils.MakePasswordHash(payload.Password)
	if err != nil {
		log.Fatal("error hashing password")
	}

	args := storage.CreateAccountParams{
		Email:    payload.Email,
		Password: hashedPassword,
	}

	_, err = accountHandler.store.CreateAccount(args)
	if err != nil {
		ResponseServerError(ctx, err)
		return
	}

	ResponseCreated(ctx)
}

func (accountHandler *AccountHandler) ListUser(ctx *gin.Context) {
	var req accountIDPayload
	if err := ctx.ShouldBindUri(&req); err != nil {
		ResponseBadRequest(ctx, err)
		return
	}

	var user models.Account
	user, err := accountHandler.store.ListAccount(req.ID)
	if err != nil {
		if errors.Is(err, storage.ErrAccountNotFound) {
			ResponseNotFound(ctx, err)
			return
		}
		ResponseServerError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, APIResponse{
		Status:  http.StatusOK,
		Message: "200 OK",
		Data:    map[string]interface{}{"user": user},
	})
}

func (accountHandler *AccountHandler) ListUsers(ctx *gin.Context) {
	users, err := accountHandler.store.ListAccounts()
	if err != nil {
		ResponseNotFound(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, APIResponse{
		Status:  http.StatusOK,
		Message: "200 OK",
		Data:    map[string]interface{}{"users": users},
	})
}

func (accountHandler *AccountHandler) UpdateUser(ctx *gin.Context) {
	var req accountIDPayload
	if err := ctx.ShouldBindUri(&req); err != nil {
		ResponseBadRequest(ctx, err)
		return
	}

	var user models.Account
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ResponseBadRequest(ctx, err)
		return
	}

	if err := accountHandler.store.UpdateAccount(&user, req.ID); err != nil {
		if errors.Is(err, storage.ErrAccountNotFound) {
			ResponseNotFound(ctx, err)
			return
		}
		ResponseServerError(ctx, err)
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}

func (accountHandler *AccountHandler) DeleteUser(ctx *gin.Context) {
	var req accountIDPayload
	if err := ctx.ShouldBindUri(&req); err != nil {
		ResponseBadRequest(ctx, err)
		return
	}

	if err := accountHandler.store.DeleteAccount(req.ID); err != nil {
		if errors.Is(err, storage.ErrAccountNotFound) {
			ResponseNotFound(ctx, err)
			return
		}
		ResponseServerError(ctx, err)
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}
