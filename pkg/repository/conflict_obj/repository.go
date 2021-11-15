package conflictObj

import (
	"errors"

	"github.com/RickardA/multiuser/pkg/aggregate"
	"github.com/google/uuid"
)

var (
	ErrConflictNotFound = errors.New("conflict not found")
)

type ConflictObjRepository interface {
	GetByID(id uuid.UUID) (aggregate.ConflictObj, error)
	Add(conflictObj aggregate.ConflictObj) error
	Delete(id uuid.UUID) error
}
