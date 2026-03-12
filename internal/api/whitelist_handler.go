package api

import (
	"context"
	"net/http"

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

	var req WhitelistRequestBody

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.Service.CreateRequest(
		context.Background(),
		req.Username,
		req.DiscordID,
		req.DiscordUsername,
		req.Message,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create request"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "request submitted",
	})
}
