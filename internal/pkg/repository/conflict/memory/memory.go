package memory_conflict_repository

import (
	"sync"

	"github.com/RickardA/multiuser/internal/pkg/domain"
	conflict_repository "github.com/RickardA/multiuser/internal/pkg/repository/conflict"
	"github.com/google/uuid"
)

type MemoryConflictObjRepository struct {
	conflicts map[uuid.UUID]domain.ConflictObj
	sync.Mutex
}

// New is a factory function to generate a new repository of customers
func New() *MemoryConflictObjRepository {
	return &MemoryConflictObjRepository{
		conflicts: make(map[uuid.UUID]domain.ConflictObj),
	}
}

// GetByID searches for a product based on it's ID
func (mpr *MemoryConflictObjRepository) GetByID(id uuid.UUID) (domain.ConflictObj, error) {
	if conflict, ok := mpr.conflicts[id]; ok {
		return conflict, nil
	}
	return domain.ConflictObj{}, conflict_repository.ErrConflictNotFound
}

// Add will add a new product to the repository
func (mpr *MemoryConflictObjRepository) Add(newConflict domain.ConflictObj) error {
	mpr.Lock()
	defer mpr.Unlock()

	mpr.conflicts[newConflict.ID] = newConflict

	return nil
}

// Update will change all values for a product based on it's ID
func (mpr *MemoryConflictObjRepository) Delete(id uuid.UUID) error {
	mpr.Lock()
	defer mpr.Unlock()

	delete(mpr.conflicts, id)
	return nil
}
