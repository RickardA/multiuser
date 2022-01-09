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
	UpdateRunway(id domain.RunwayID, input domain.Runway) (domain.Runway, error)
	DeleteRunwayWithID(id domain.RunwayID) error
}

type ConflictRepository interface {
	GetConflictForRunway(designator string) (domain.ConflictObj, error)
	CreateConflictObj(input domain.ConflictObj) (domain.ConflictObj, error)
	UpdateConflictObj(input domain.ConflictObj) (domain.ConflictObj, error)
	DeleteConflictWithID(id string) error
}
