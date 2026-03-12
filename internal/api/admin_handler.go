package api

import (
	"context"
	"net/http"

	"mc-webserver/internal/service"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	Service          *service.AdminService
	WhitelistService *service.WhitelistService
}

func NewAdminHandler(
	adminService *service.AdminService,
	whitelistService *service.WhitelistService,
) *AdminHandler {

	return &AdminHandler{
		Service:          adminService,
		WhitelistService: whitelistService,
	}
}

type AdminLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *AdminHandler) Login(c *gin.Context) {

	var req AdminLoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ok, err := h.Service.Login(
		context.Background(),
		req.Username,
		req.Password,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}

	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func (h *AdminHandler) ApprovePlayer(c *gin.Context) {

	username := c.Param("username")

	err := h.WhitelistService.ApprovePlayer(
		context.Background(),
		username,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not approve"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "player whitelisted",
	})
}
