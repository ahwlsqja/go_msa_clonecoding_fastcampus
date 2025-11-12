package main

import (
	"userfc/cmd/user/repository"
	"userfc/cmd/user/resource"
	"userfc/config"
	"userfc/handler"
	"userfc/infrastructure/log"
	"userfc/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()
	redis := resource.InitRedis(&cfg)
	db := resource.InitDB(&cfg)

	log.SetupLogger()

	userRepository := repository.NewUserRepository(db, redis)
	userHandler := handler.NewUserHandler()

	port := cfg.App.Port
	router := gin.Default()

	// Setup routes
	routes.SetupRoutes(router, *userHandler)

	log.Logger.Printf("Server starting on port: %s", port)
	router.Run(":" + port)
}