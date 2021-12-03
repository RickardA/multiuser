package runway

import (
	"errors"

	"github.com/RickardA/multiuser/internal/pkg/domain"
)

var (
	ErrRunwayNotFound = errors.New("the runway was not found")
)

type RunwayRepository interface {
	GetByDesignator(designator string) (domain.Runway, error)
	Add(runway domain.Runway) error
	Update(runway domain.Runway) error
}
