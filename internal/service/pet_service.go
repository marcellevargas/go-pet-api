package service

import (
	"go-pet-api/internal/domain"
	"go-pet-api/internal/repository"

	"github.com/google/uuid"
)

type PetService struct {
	repo repository.PetRepository
}

func NewPetService(repo repository.PetRepository) *PetService {
	return &PetService{
		repo: repo,
	}
}

func (s *PetService) CreatePet(pet *domain.Pet) error {
	pet.ID = uuid.New().String()
	return s.repo.Create(pet)
}

func (s *PetService) GetPet(id string) (*domain.Pet, error) {
	return s.repo.GetByID(id)
}

func (s *PetService) GetAllPets() ([]*domain.Pet, error) {
	return s.repo.GetAll()
}

func (s *PetService) UpdatePet(pet *domain.Pet) error {
	return s.repo.Update(pet)
}

func (s *PetService) DeletePet(id string) error {
	return s.repo.Delete(id)
}
