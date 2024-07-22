package main

import (
	"bae-backend/internal/adapter/config"
	"bae-backend/internal/adapter/handler/http"
	"bae-backend/internal/adapter/handler/http/huser"
	"bae-backend/internal/adapter/storage/mongodb"
	"bae-backend/internal/adapter/storage/mongodb/repository"
	"bae-backend/internal/core/service"
	"fmt"
	"log/slog"
	"os"
)

func main() {
	// Load environment variables
	config, err := config.New()
	if err != nil {
		slog.Error("Error loading environment variables", "error", err)
		os.Exit(1)
	}

	// Init database
	db, err := mongodb.New(config.DB)
	if err != nil {
		slog.Error("Error initializing database connection", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	// Dependency injection
	// User
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := huser.NewUserHandler(userService)

	// Init router
	router, err := http.NewRouter(
		config.HTTP,
		userHandler,
	)
	if err != nil {
		slog.Error("Error initializing router", "error", err)
		os.Exit(1)
	}

	// Start server
	listenAddr := fmt.Sprintf("%s:%s", config.HTTP.URL, config.HTTP.Port)
	slog.Info("Starting the HTTP server", "listen_address", listenAddr)
	err = router.Serve(listenAddr)
	if err != nil {
		slog.Error("Error starting the HTTP server", "error", err)
		os.Exit(1)
	}

}
