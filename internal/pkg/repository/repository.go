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
	GetRunwayByID(id string) (domain.Runway, error)
	CreateRunway(domain.Runway) (domain.Runway, error)
	UpdateRunway(domain.Runway) (domain.Runway, error)
	DeleteRunwayWithID(id string) error
}

type ConflictRepository interface {
	GetConflictForRunway(designator string) (domain.ConflictObj, error)
	CreateConflictObj(domain.ConflictObj) (domain.ConflictObj, error)
	UpdateConflictObj(domain.ConflictObj) (domain.ConflictObj, error)
	DeleteConflictWithID(id string) error
}
