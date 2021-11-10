package aggregate

import (
	"errors"

	"github.com/RickardA/multiuser/pkg/entity"
	"github.com/google/uuid"
)

var (
	ErrMissingZoneIdentity = errors.New("Missing zone identity")
)

type Zone struct {
	ID            uuid.UUID            `json:"id"`
	Identity      string               `json:"identity"`
	Contamination *entity.IntegerField `json:"contamination"`
	Depth         *entity.IntegerField `json:"depth"`
	Coverage      *entity.IntegerField `json:"coverage"`
}

type ZoneComparisionObj struct {
	ID            uuid.UUID                          `json:"-"`
	Identity      string                             `json:"identity"`
	Contamination *entity.IntegerFieldComparisionObj `json:"contamination"`
	Depth         *entity.IntegerFieldComparisionObj `json:"depth"`
	Coverage      *entity.IntegerFieldComparisionObj `json:"coverage"`
}

func CreateZone(identity string) (Zone, error) {
	if identity == "" {
		return Zone{}, ErrMissingZoneIdentity
	}

	return Zone{
		ID:            uuid.New(),
		Identity:      identity,
		Contamination: &entity.IntegerField{},
		Depth:         &entity.IntegerField{},
		Coverage:      &entity.IntegerField{},
	}, nil
}
