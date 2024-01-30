package handlers

import "github.com/gin-gonic/gin"

type PingHandler struct{}

func NewPingHandlers() *PingHandler {
	return &PingHandler{}
}

func (pingHandler *PingHandler) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
