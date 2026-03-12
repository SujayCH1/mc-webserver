package api

import (
	"context"

	"mc-webserver/internal/service"

	"github.com/gin-gonic/gin"
)

type WhitelistHandler struct {
	Service *service.WhitelistService
}

func NewWhitelistHandler(s *service.WhitelistService) *WhitelistHandler {
	return &WhitelistHandler{Service: s}
}

type WhitelistRequestBody struct {
	Username        string `json:"username"`
	DiscordID       string `json:"discord_id"`
	DiscordUsername string `json:"discord_username"`
	Message         string `json:"message"`
}

func (h *WhitelistHandler) CreateRequest(c *gin.Context) {

	username := c.GetString("username")

	var req struct {
		Message string `json:"message"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := h.Service.CreateRequest(
		context.Background(),
		username,
		req.Message,
	)

	if err != nil {
		c.JSON(500, gin.H{"error": "could not create request"})
		return
	}

	c.JSON(200, gin.H{
		"message": "request submitted",
	})
}
