package repository

import (
	"errors"

	"github.com/RickardA/multiuser/internal/pkg/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	GetRunwayByID(id primitive.ObjectID) (domain.Runway, error)
	CreateRunway(input domain.Runway) (primitive.ObjectID, error)
	UpdateRunway(input domain.Runway) (domain.Runway, error)
	DeleteRunwayWithID(id string) error
}

type ConflictRepository interface {
	GetConflictForRunway(designator string) (domain.ConflictObj, error)
	CreateConflictObj(input domain.ConflictObj) (domain.ConflictObj, error)
	UpdateConflictObj(input domain.ConflictObj) (domain.ConflictObj, error)
	DeleteConflictWithID(id string) error
}
