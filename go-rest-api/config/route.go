package config

import (
	"go-rest-api/internal"
	"go-rest-api/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(app *internal.App) *gin.Engine {
	r := gin.Default()

	// Public route
	r.GET("/login", app.Handler.Login)

	// Protected group
	auth := r.Group("/", middleware.JWTAuth())
	{
		auth.GET("/employees", app.Handler.GetEmployees)
	}

	return r
}
