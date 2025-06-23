package handlers

import (
	"go-pet-api/internal/domain"
	"go-pet-api/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PetHandler struct {
	service *service.PetService
}

func NewPetHandler(service *service.PetService) *PetHandler {
	return &PetHandler{
		service: service,
	}
}

func (h *PetHandler) CreatePet(c *gin.Context) {
	var pet domain.Pet
	if err := c.ShouldBindJSON(&pet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := h.service.CreatePet(&pet); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, pet)
}

func (h *PetHandler) GetPet(c *gin.Context) {
	id := c.Param("id")

	pet, err := h.service.GetPet(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pet not found"})
		return
	}

	c.JSON(http.StatusOK, pet)
}

func (h *PetHandler) GetAllPets(c *gin.Context) {
	pets, err := h.service.GetAllPets()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, pets)
}

func (h *PetHandler) UpdatePet(c *gin.Context) {
	id := c.Param("id")

	var pet domain.Pet
	if err := c.ShouldBindJSON(&pet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	pet.ID = id
	if err := h.service.UpdatePet(&pet); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, pet)
}

func (h *PetHandler) DeletePet(c *gin.Context) {
	id := c.Param("id")

	if err := h.service.DeletePet(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
