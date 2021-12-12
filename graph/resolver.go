package graph

import "github.com/RickardA/multiuser/internal/pkg/domain"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	RunwayDB runwayRepository
}

type runwayRepository interface {
	GetByDesignator(designator string) (domain.Runway, error)
	Add(runway domain.Runway) error
	Update(runway domain.Runway) error
}
