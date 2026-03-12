package api

import (
	"context"
	"mc-webserver/internal/service"

	"github.com/gin-gonic/gin"
)

type ServerAuthHandler struct {
	Service *service.AuthService
}

func NewServerAuthHandler(s *service.AuthService) *ServerAuthHandler {
	return &ServerAuthHandler{Service: s}
}

func (h *ServerAuthHandler) CheckPlayer(c *gin.Context) {

	var req struct {
		Username string `json:"username"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	player, err := h.Service.GetPlayer(context.Background(), req.Username)

	if err != nil {
		c.JSON(404, gin.H{"exists": false})
		return
	}

	c.JSON(200, gin.H{
		"exists": true,
		"banned": player.Banned,
	})
}

func (h *ServerAuthHandler) Authenticate(c *gin.Context) {

	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	player, err := h.Service.GetPlayer(c.Request.Context(), req.Username)

	if err != nil {
		c.JSON(401, gin.H{"allowed": false})
		return
	}

	if player.Banned {
		c.JSON(403, gin.H{"allowed": false})
		return
	}

	ok, err := h.Service.VerifyPassword(
		c.Request.Context(),
		req.Username,
		req.Password,
	)

	if err != nil || !ok {
		c.JSON(401, gin.H{
			"allowed": false,
		})
		return
	}

	c.JSON(200, gin.H{
		"allowed": true,
	})
}
