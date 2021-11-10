package runway

import (
	"errors"

	"github.com/RickardA/multiuser/pkg/aggregate"
)

var (
	ErrRunwayNotFound = errors.New("the runway was not found")
)

type RunwayRepository interface {
	GetByDesignator(designator string) (aggregate.Runway, error)
	Add(runway aggregate.Runway) error
	Update(runway aggregate.Runway) error
}
