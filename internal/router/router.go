package router

import (
	"database/sql"
	"time"

	"mc-webserver/internal/api"
	"mc-webserver/internal/middleware"
	"mc-webserver/internal/repository"
	"mc-webserver/internal/service"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
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

	r.POST("/register",
		middleware.RateLimit(rate.Every(time.Minute/3), 1),
		authHandler.Register,
	)

	r.POST("/login",
		middleware.RateLimit(rate.Every(time.Minute/5), 1),
		authHandler.Login,
	)

	r.POST("/whitelist/request",
		middleware.JWTAuthMiddleware(),
		middleware.RateLimit(rate.Every(time.Minute/2), 1),
		whitelistHandler.CreateRequest,
	)

	admin := r.Group("/admin")
	admin.Use(middleware.AdminOnly())
	{
		admin.GET("/requests", adminHandler.GetRequests)
		admin.POST("/approve/:username", adminHandler.ApprovePlayer)
	}

	r.POST("/admin/approve/:username", adminHandler.ApprovePlayer)

	return r
}
