// Package memory is a in memory implementation of the ProductRepository interface.
package memory

import (
	"sync"

	"github.com/RickardA/multiuser/pkg/aggregate"
	"github.com/RickardA/multiuser/pkg/repository/runway"
)

type MemoryRunwayRepository struct {
	runways map[string]aggregate.Runway
	sync.Mutex
}

// New is a factory function to generate a new repository of customers
func New() *MemoryRunwayRepository {
	return &MemoryRunwayRepository{
		runways: make(map[string]aggregate.Runway),
	}
}

// GetByID searches for a product based on it's ID
func (mpr *MemoryRunwayRepository) GetByDesignator(designator string) (aggregate.Runway, error) {
	if runway, ok := mpr.runways[designator]; ok {
		return runway, nil
	}
	return aggregate.Runway{}, runway.ErrRunwayNotFound
}

// Add will add a new product to the repository
func (mpr *MemoryRunwayRepository) Add(newRwy aggregate.Runway) error {
	mpr.Lock()
	defer mpr.Unlock()

	mpr.runways[newRwy.Designator] = newRwy

	return nil
}

// Update will change all values for a product based on it's ID
func (mpr *MemoryRunwayRepository) Update(updRwy aggregate.Runway) error {
	mpr.Lock()
	defer mpr.Unlock()

	if _, ok := mpr.runways[updRwy.Designator]; !ok {
		return runway.ErrRunwayNotFound
	}

	mpr.runways[updRwy.Designator] = updRwy
	return nil
}
