package internal

import (
	"database/sql"
	"net/http"

	"go-rest-api/internal/employee"
	"go-rest-api/internal/pokemon"
)

type App struct {
	Mux      *http.ServeMux
	Employee *employee.Handler
	Pokemon  *pokemon.Handler
}

func Init(db *sql.DB) *App {
	repo := employee.NewRepository(db)
	service := employee.NewService(repo)
	handler := employee.NewHandler(service)
	pokemonRepo := pokemon.NewRepository()
	pokemonService := pokemon.NewService(pokemonRepo)
	pokemonHandler := pokemon.NewHandler(pokemonService)

	return &App{
		Mux:      http.NewServeMux(), // Routing pindah ke config/server.go
		Employee: handler,
		Pokemon:  pokemonHandler,
	}
}
