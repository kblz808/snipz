package router

import (
	"log/slog"
	"snipz/internal/http/handlers"
	"snipz/internal/http/middlewares"
	"snipz/internal/services"
	"snipz/internal/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"
)

type Router struct {
	*gin.Engine
}

func NewRouter(
	config utils.Config,
	userHandler handlers.UserHandler,
	authHandler handlers.AuthHandler,
	tokenService services.TokenService,
) (*Router, error) {
	ginConfig := cors.DefaultConfig()
	// allowedOrigins := string[]{}
	// ginConfig.AllowOrigins = allowedOrigins

	router := gin.New()
	router.Use(sloggin.New(slog.Default()), gin.Recovery(), cors.New(ginConfig))

	v1 := router.Group("/v1")
	{
		user := v1.Group("/users")
		{
			user.POST("/", userHandler.Register)
			user.POST("/login", authHandler.Login)

			_ = user.Group("/").Use(middlewares.AuthMiddleware(tokenService))
			{
				// authUser.GET("/", userHandler.GetAll)
				// authUser.GET("/:id", userHandler.Get)
				// authUser.PUT("/:id", userHandler.Update)
			}
		}
	}

	return nil, nil
}
