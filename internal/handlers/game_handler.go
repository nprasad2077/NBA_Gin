package handlers

import (
	"net/http"
	"github.com/nprasad2077/NBA_Gin/internal/domain"
	"github.com/gin-gonic/gin"
)

type GameHandler struct {
	Repo domain.GameRepository
}

func (h *GameHandler) GetGame(c *gin.Context) {
	id := c.Param("id")
	game, err := h.Repo.GetGameByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Game not found"})
		return
	}
	c.JSON(http.StatusOK, game)
}

func (h *GameHandler) ListGames(c *gin.Context) {
	var filter domain.GameFilter

	// 1. Bind Query Params to Struct
    // This handles ?page=1, ?team=LAL, ?season=2015, etc.
	if err := c.ShouldBindQuery(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameters"})
		return
	}

	// 2. Call Repository
	response, err := h.Repo.ListGames(c.Request.Context(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}