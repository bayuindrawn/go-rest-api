package config

import (
	"fmt"
	"time"

	"github.com/didip/tollbooth/v7"
	tollbooth_gin "github.com/didip/tollbooth_gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"go-rest-api/internal"
	"go-rest-api/internal/middleware"
)

func SetupRoutes(app *internal.App) *gin.Engine {
	r := gin.New()

	// Custom Logging
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("[%s] %s %s %d %s \"%s\"\n",
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
		)
	}))
	r.Use(gin.Recovery())

	// CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Rate limiting menggunakan memory counter sederhana
	limiter := tollbooth.NewLimiter(5, nil)
	authLimiter := tollbooth_gin.LimitHandler(limiter)

	r.POST("/login", authLimiter, app.Handler.Login)
	r.POST("/refresh", authLimiter, app.Handler.Refresh)

	auth := r.Group("/")
	auth.Use(middleware.JWTAuth())
	auth.Use(authLimiter)
	{
		auth.GET("/employees", app.Handler.GetEmployees)
	}

	return r
}
