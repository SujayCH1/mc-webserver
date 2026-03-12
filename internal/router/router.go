package router

import (
	"database/sql"

	"mc-webserver/internal/api"
	"mc-webserver/internal/repository"
	"mc-webserver/internal/service"

	"github.com/gin-gonic/gin"
)

func SetUpRouter(db *sql.DB) *gin.Engine {

	r := gin.Default()

	playerRepo := repository.NewPlayerRepository(db)
	requestRepo := repository.NewWhitelistRequestRepository(db)
	adminRepo := repository.NewAdminRepository(db)

	authService := service.NewAuthService(playerRepo)
	whitelistService := service.NewWhitelistService(playerRepo, requestRepo)
	adminService := service.NewAdminService(adminRepo)

	authHandler := api.NewAuthHandler(authService)
	whitelistHandler := api.NewWhitelistHandler(whitelistService)
	adminHandler := api.NewAdminHandler(adminService, whitelistService)

	r.POST("/register", authHandler.Register)
	r.POST("/login", authHandler.Login)

	r.POST("/whitelist/request", whitelistHandler.CreateRequest)

	r.POST("/admin/login", adminHandler.Login)
	r.POST("/admin/approve/:username", adminHandler.ApprovePlayer)

	return r
}
