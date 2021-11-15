// Package memory is a in memory implementation of the ProductRepository interface.
package memory_conflictObj

import (
	"sync"

	conflictObj "github.com/RickardA/multiuser/pkg/repository/conflict_obj"

	"github.com/RickardA/multiuser/pkg/aggregate"
	"github.com/google/uuid"
)

type MemoryConflictObjRepository struct {
	conflicts map[uuid.UUID]aggregate.ConflictObj
	sync.Mutex
}

// New is a factory function to generate a new repository of customers
func New() *MemoryConflictObjRepository {
	return &MemoryConflictObjRepository{
		conflicts: make(map[uuid.UUID]aggregate.ConflictObj),
	}
}

// GetByID searches for a product based on it's ID
func (mpr *MemoryConflictObjRepository) GetByID(id uuid.UUID) (aggregate.ConflictObj, error) {
	if conflict, ok := mpr.conflicts[id]; ok {
		return conflict, nil
	}
	return aggregate.ConflictObj{}, conflictObj.ErrConflictNotFound
}

// Add will add a new product to the repository
func (mpr *MemoryConflictObjRepository) Add(newConflict aggregate.ConflictObj) error {
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
