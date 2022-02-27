package repository

import (
	"errors"

	"github.com/RickardA/multiuser/internal/pkg/domain"
)

var (
	ErrNotImplemented = errors.New("function not implemented")
)

type Repository interface {
	RunwayRepository
	ConflictRepository
}

type RunwayRepository interface {
	GetRunwayByDesignator(designator string) (domain.Runway, error)
	GetRunwayByID(id domain.RunwayID) (domain.Runway, error)
	CreateRunway(input domain.Runway) (domain.RunwayID, error)
	UpdateRunway(id domain.RunwayID, input domain.Runway, clientID string) (domain.Runway, error)
	DeleteRunwayWithID(id domain.RunwayID) error
}

type ConflictRepository interface {
	GetConflictByID(id domain.ConflictID) (domain.Conflict, error)
	GetConflictForRunway(runwayID domain.RunwayID) (domain.Conflict, error)
	CreateConflict(input domain.Conflict) (domain.ConflictID, error)
	UpdateConflict(input domain.Conflict) (domain.Conflict, error)
	DeleteConflictWithID(id domain.ConflictID) error
}
