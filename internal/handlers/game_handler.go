package handlers

import (
	"net/http"
	"strconv"
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
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))
	offset := (page - 1) * pageSize

	games, err := h.Repo.ListGames(c.Request.Context(), pageSize, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, games)
}