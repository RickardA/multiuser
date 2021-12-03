// Package memory is a in memory implementation of the ProductRepository interface.
package memory_runway

import (
	"sync"

	"github.com/RickardA/multiuser/internal/pkg/domain"
	"github.com/RickardA/multiuser/internal/pkg/repository/runway"
)

type MemoryRunwayRepository struct {
	runways map[string]domain.Runway
	sync.Mutex
}

// New is a factory function to generate a new repository of customers
func New() *MemoryRunwayRepository {
	return &MemoryRunwayRepository{
		runways: make(map[string]domain.Runway),
	}
}

// GetByID searches for a product based on it's ID
func (mpr *MemoryRunwayRepository) GetByDesignator(designator string) (domain.Runway, error) {
	if runway, ok := mpr.runways[designator]; ok {
		return runway, nil
	}
	return domain.Runway{}, runway.ErrRunwayNotFound
}

// Add will add a new product to the repository
func (mpr *MemoryRunwayRepository) Add(newRwy domain.Runway) error {
	mpr.Lock()
	defer mpr.Unlock()

	mpr.runways[newRwy.Designator] = newRwy

	return nil
}

// Update will change all values for a product based on it's ID
func (mpr *MemoryRunwayRepository) Update(updRwy domain.Runway) error {
	mpr.Lock()
	defer mpr.Unlock()

	if _, ok := mpr.runways[updRwy.Designator]; !ok {
		return runway.ErrRunwayNotFound
	}

	mpr.runways[updRwy.Designator] = updRwy
	return nil
}
