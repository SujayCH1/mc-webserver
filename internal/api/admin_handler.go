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

	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ok, err := h.Service.Login(
		context.Background(),
		req.Username,
		req.Password,
	)

	if err != nil || !ok {
		c.JSON(401, gin.H{"error": "invalid credentials"})
		return
	}

	token, err := h.Service.GenerateAdminToken(req.Username)
	if err != nil {
		c.JSON(500, gin.H{"error": "token error"})
		return
	}

	c.JSON(200, gin.H{
		"token": token,
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

func (h *AdminHandler) GetRequests(c *gin.Context) {

	requests, err := h.WhitelistService.GetPendingRequests(context.Background())
	if err != nil {
		c.JSON(500, gin.H{
			"error": "could not fetch requests",
		})
		return
	}

	c.JSON(200, requests)
}
