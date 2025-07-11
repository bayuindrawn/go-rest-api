package employee

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetEmployees(w http.ResponseWriter, r *http.Request) {
	employees, err := h.service.GetEmployees()
	if err != nil {
		http.Error(w, "Failed to get employees", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(employees)
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	password := r.URL.Query().Get("password")

	if username != "admin" || password != "1234" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(2 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("JWT_SECRET")
	tokenStr, err := token.SignedString([]byte(secret))
	if err != nil {
		http.Error(w, "Failed to sign token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": tokenStr})
}
