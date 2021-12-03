package conflict_repository

import (
	"errors"

	"github.com/RickardA/multiuser/internal/pkg/domain"
	"github.com/google/uuid"
)

var (
	ErrConflictNotFound = errors.New("conflict not found")
)

type ConflictObjRepository interface {
	GetByID(id uuid.UUID) (domain.ConflictObj, error)
	Add(conflictObj domain.ConflictObj) error
	Delete(id uuid.UUID) error
}
