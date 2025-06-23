package main

import (
	"go-pet-api/internal/handlers"
	"go-pet-api/internal/repository"
	"go-pet-api/internal/service"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize repository
	repo := repository.NewInMemoryPetRepository()

	// Initialize service
	petService := service.NewPetService(repo)

	// Initialize handler
	petHandler := handlers.NewPetHandler(petService)

	// Initialize Gin router
	router := gin.Default()

	// Register routes
	router.POST("/pets", petHandler.CreatePet)
	router.GET("/pets", petHandler.GetAllPets)
	router.GET("/pets/:id", petHandler.GetPet)
	router.PUT("/pets/:id", petHandler.UpdatePet)
	router.DELETE("/pets/:id", petHandler.DeletePet)

	// Start server
	log.Println("Server starting on port 8080...")
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
