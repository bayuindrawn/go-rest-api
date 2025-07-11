package internal

import (
	"database/sql"
	"net/http"

	"go-rest-api/internal/employee"
)

type App struct {
	Mux     *http.ServeMux
	Handler *employee.Handler
}

func Init(db *sql.DB) *App {
	repo := employee.NewRepository(db)
	service := employee.NewService(repo)
	handler := employee.NewHandler(service)

	return &App{
		Mux:     http.NewServeMux(), // Routing pindah ke config/server.go
		Handler: handler,
	}
}
