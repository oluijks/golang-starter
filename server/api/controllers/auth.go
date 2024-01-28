package controllers

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/oluijks/golang-starter/server/api/models"
	"github.com/oluijks/golang-starter/server/internal/config"
	"github.com/oluijks/golang-starter/server/internal/storage"
	"github.com/oluijks/golang-starter/server/internal/token"
	"github.com/oluijks/golang-starter/server/internal/utils"
)

type AuthHandler struct {
	store      storage.Store
	config     *config.Config
	tokenMaker token.Maker
}

func NewAuthHandlers(store storage.Store, config *config.Config) *AuthHandler {
	tokenMaker, err := token.PasetoMaker(config.SymmectricTokenKey)
	if err != nil {
		log.Fatal("Error creating token maker:", err)
	}
	return &AuthHandler{store: store, config: config, tokenMaker: tokenMaker}
}

type signInPayload struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type accountResponse struct {
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func newAccountResponse(account models.Account) accountResponse {
	return accountResponse{
		Email:     account.Email,
		CreatedAt: account.CreatedAt,
		UpdatedAt: account.UpdatedAt,
	}
}

type signInReponse struct {
	Account     accountResponse `json:"account"`
	AccessToken string          `json:"access_token"`
}

func (authHandler *AuthHandler) SignIn(ctx *gin.Context) {
	var payload signInPayload
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ResponseBadRequest(ctx, err)
		return
	}

	account, err := authHandler.store.GetAccountByEmail(payload.Email)
	if err != nil {
		if errors.Is(err, storage.ErrAccountNotFound) {
			ResponseNotFound(ctx, err)
			return
		}
		ResponseServerError(ctx, err)
		return
	}

	err = utils.ComparePassword(account.Password, payload.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{})
		return
	}

	accessToken, err := authHandler.tokenMaker.MakeToken(account.Email, authHandler.config.AccessTokenDuration)
	if err != nil {
		ResponseServerError(ctx, err)
		return
	}

	authResponse := signInReponse{
		Account:     newAccountResponse(account),
		AccessToken: accessToken,
	}

	ctx.JSON(http.StatusOK, APIResponse{
		Status:  http.StatusOK,
		Message: "200 OK",
		Data:    map[string]interface{}{"auth": authResponse},
	})
}
