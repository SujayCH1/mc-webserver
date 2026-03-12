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

	// Serve React static files
	r.Static("/assets", "./frontend/dist/assets")

	r.GET("/", func(c *gin.Context) {
		c.File("./frontend/dist/index.html")
	})

	// Repositories
	playerRepo := repository.NewPlayerRepository(db)
	requestRepo := repository.NewWhitelistRequestRepository(db)
	adminRepo := repository.NewAdminRepository(db)

	// Services
	authService := service.NewAuthService(playerRepo)
	whitelistService := service.NewWhitelistService(playerRepo, requestRepo)
	adminService := service.NewAdminService(adminRepo)

	// Handlers
	authHandler := api.NewAuthHandler(authService)
	serverAuthHandler := api.NewServerAuthHandler(authService)
	whitelistHandler := api.NewWhitelistHandler(whitelistService)
	adminHandler := api.NewAdminHandler(adminService, whitelistService)

	// Public APIs
	r.POST("/register",
		middleware.RateLimit(rate.Every(time.Minute/3), 1),
		authHandler.Register,
	)

	r.POST("/login",
		middleware.RateLimit(rate.Every(time.Minute/5), 1),
		authHandler.Login,
	)

	// Player APIs (JWT protected)
	r.POST("/whitelist/request",
		middleware.JWTAuthMiddleware(),
		middleware.RateLimit(rate.Every(time.Minute/2), 1),
		whitelistHandler.CreateRequest,
	)

	// Admin APIs
	admin := r.Group("/admin")
	admin.Use(middleware.AdminOnly())
	{
		admin.GET("/requests", adminHandler.GetRequests)
		admin.POST("/approve/:username", adminHandler.ApprovePlayer)
	}

	// Minecraft Server APIs
	serverAuth := r.Group("/server")
	serverAuth.Use(middleware.ServerOnly())
	{
		serverAuth.POST("/check-player", serverAuthHandler.CheckPlayer)
		serverAuth.POST("/auth", serverAuthHandler.Authenticate)
	}

	// React fallback for SPA routing
	r.NoRoute(func(c *gin.Context) {
		c.File("./frontend/dist/index.html")
	})

	return r
}
