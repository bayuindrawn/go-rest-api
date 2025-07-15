package pokemon

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetPokemons(c *gin.Context) {
	ctx := c.Request.Context()

	data, err := h.service.FetchPokemonList(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch pokemons"})
		return
	}
	c.JSON(http.StatusOK, data)
}
