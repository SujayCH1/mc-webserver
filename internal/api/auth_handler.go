package api

import (
	"context"
	"net/http"

	"mc-webserver/internal/service"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	Service *service.AuthService
}

func NewAuthHandler(s *service.AuthService) *AuthHandler {
	return &AuthHandler{Service: s}
}

type RegisterRequest struct {
	Username        string `json:"username"`
	DiscordID       string `json:"discord_id"`
	DiscordUsername string `json:"discord_username"`
	Password        string `json:"password"`
}

func (h *AuthHandler) Register(c *gin.Context) {

	var req RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.Service.Register(
		context.Background(),
		req.Username,
		req.DiscordID,
		req.DiscordUsername,
		req.Password,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "player registered",
	})
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *AuthHandler) Login(c *gin.Context) {

	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ok, err := h.Service.VerifyPassword(
		context.Background(),
		req.Username,
		req.Password,
	)

	if err != nil || !ok {
		c.JSON(401, gin.H{"error": "invalid credentials"})
		return
	}

	token, err := h.Service.GenerateToken(req.Username, "player")
	if err != nil {
		c.JSON(500, gin.H{"error": "token error"})
		return
	}

	c.JSON(200, gin.H{
		"token": token,
	})
}
