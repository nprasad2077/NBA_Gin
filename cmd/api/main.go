package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nprasad2077/NBA_Gin/internal/handlers"
	"github.com/nprasad2077/NBA_Gin/internal/repository"
	
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	// 1. Config Setup
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		slog.Warn("No .env file found, relying on environment variables")
	}

	// 2. Database Setup
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		viper.GetString("DB_USER"),
		viper.GetString("DB_PASSWORD"),
		viper.GetString("DB_HOST"),
		viper.GetString("DB_PORT"),
		viper.GetString("DB_NAME"),
	)
	
	db, err := repository.NewPostgresDB(dsn)
	if err != nil {
		slog.Error("Failed to initialize DB", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	// 3. Init Layers
	gameRepo := repository.NewGameRepository(db)
	gameHandler := &handlers.GameHandler{Repo: gameRepo}

	// 4. Router Setup
	r := gin.Default()
	
	// Health Check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	api := r.Group("/api/v1")
	{
		api.GET("/games", gameHandler.ListGames)
		api.GET("/games/:id", gameHandler.GetGame)
	}

	// 5. Server Start with Graceful Shutdown
	srv := &http.Server{
		Addr:    ":" + viper.GetString("PORT"),
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	slog.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		slog.Error("Server forced to shutdown: ", err)
	}
	slog.Info("Server exiting")
}