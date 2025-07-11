package config

import (
	"net/http"

	"go-rest-api/internal"
	"go-rest-api/internal/middleware"
)

func SetupRoutes(app *internal.App) *http.ServeMux {
	mux := http.NewServeMux()

	// Public routes
	mux.HandleFunc("/login", app.Handler.Login)

	// Protected routes
	mux.Handle("/employees", middleware.JWTAuth(http.HandlerFunc(app.Handler.GetEmployees)))

	return mux
}
