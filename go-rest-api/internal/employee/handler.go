package employee

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Handler struct {
	service Service
}

func NewHandler(s Service) *Handler {
	return &Handler{service: s}
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.Username != "admin" || req.Password != "1234" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	jwtSecret := []byte(os.Getenv("JWT_SECRET"))
	refreshSecret := []byte(os.Getenv("REFRESH_SECRET"))

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": req.Username,
		"exp":      time.Now().Add(15 * time.Minute).Unix(),
	})
	atStr, _ := accessToken.SignedString(jwtSecret)

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": req.Username,
		"exp":      time.Now().Add(7 * 24 * time.Hour).Unix(),
	})
	rtStr, _ := refreshToken.SignedString(refreshSecret)

	c.JSON(http.StatusOK, gin.H{
		"access_token":  atStr,
		"refresh_token": rtStr,
		"expires_in":    15 * 60,
	})
}

func (h *Handler) Refresh(c *gin.Context) {
	var body struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := jwt.Parse(body.RefreshToken, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("REFRESH_SECRET")), nil
	})
	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		return
	}

	claims := token.Claims.(jwt.MapClaims)
	username := claims["username"].(string)

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(15 * time.Minute).Unix(),
	})
	atStr, _ := accessToken.SignedString([]byte(os.Getenv("JWT_SECRET")))

	c.JSON(http.StatusOK, gin.H{"access_token": atStr, "expires_in": 15 * 60})
}

func (h *Handler) GetEmployees(c *gin.Context) {
	employees, err := h.service.GetEmployees()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get employees"})
		return
	}
	c.JSON(http.StatusOK, employees)
}

func (h *Handler) GetCounter(c *gin.Context) {
	h.service.IncrementCounter()
	count := h.service.GetCounter()

	c.JSON(200, gin.H{
		"message": "Counter accessed",
		"count":   count,
	})
}
