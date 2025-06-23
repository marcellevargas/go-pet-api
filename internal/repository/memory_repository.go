package repository

import (
	"errors"
	"go-pet-api/internal/domain"
	"sync"
)

type InMemoryPetRepository struct {
	pets  map[string]*domain.Pet
	mutex sync.RWMutex
}

func NewInMemoryPetRepository() *InMemoryPetRepository {
	return &InMemoryPetRepository{
		pets: make(map[string]*domain.Pet),
	}
}

func (r *InMemoryPetRepository) Create(pet *domain.Pet) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.pets[pet.ID]; exists {
		return errors.New("pet already exists")
	}

	r.pets[pet.ID] = pet
	return nil
}

func (r *InMemoryPetRepository) GetByID(id string) (*domain.Pet, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	pet, exists := r.pets[id]
	if !exists {
		return nil, errors.New("pet not found")
	}

	return pet, nil
}

func (r *InMemoryPetRepository) GetAll() ([]*domain.Pet, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	pets := make([]*domain.Pet, 0, len(r.pets))
	for _, pet := range r.pets {
		pets = append(pets, pet)
	}

	return pets, nil
}

func (r *InMemoryPetRepository) Update(pet *domain.Pet) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.pets[pet.ID]; !exists {
		return errors.New("pet not found")
	}

	r.pets[pet.ID] = pet
	return nil
}

func (r *InMemoryPetRepository) Delete(id string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.pets[id]; !exists {
		return errors.New("pet not found")
	}

	delete(r.pets, id)
	return nil
}
