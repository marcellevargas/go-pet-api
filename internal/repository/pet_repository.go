package repository

import "go-pet-api/internal/domain"

type PetRepository interface {
	Create(pet *domain.Pet) error
	GetByID(id string) (*domain.Pet, error)
	GetAll() ([]*domain.Pet, error)
	Update(pet *domain.Pet) error
	Delete(id string) error
}
