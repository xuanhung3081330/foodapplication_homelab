package main

import (
	"log"
	"foodapplication/config"
	"foodapplication/handlers"
	repoImp "foodapplication/repositories/implementations"
	svcImp "foodapplication/services/implementations"
	"github.com/gin-gonic/gin"
)

func main(){
	// Connect to DB
	if err := config.ConnectDB(); err != nil {
		log.Fatalf("DB Connection failed: %v", err)
	}

	// Setup dependencies
	repo := repoImp.NewFoodRepository()
	service := svcImp.NewFoodService(repo)
	handler := handlers.NewFoodHandler(service)

	// Start Gin
	router := gin.Default()
	router.POST("/foods", handler.CreateFood)
	router.Run("localhost:8080")
}
