package domain

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrMissingZoneIdentity = errors.New("Missing zone identity")
)

type Zone struct {
	ID            uuid.UUID     `json:"id"`
	Identity      string        `json:"identity"`
	Contamination *IntegerField `diff:"contamination"`
	Depth         *IntegerField `diff:"depth"`
	Coverage      *IntegerField `diff:"coverage"`
}

type ZoneComparisionObj struct {
	ID            uuid.UUID                   `json:"-"`
	Identity      string                      `json:"identity"`
	Contamination *IntegerFieldComparisionObj `json:"contamination"`
	Depth         *IntegerFieldComparisionObj `json:"depth"`
	Coverage      *IntegerFieldComparisionObj `json:"coverage"`
}

func CreateZone(identity string) (Zone, error) {
	if identity == "" {
		return Zone{}, ErrMissingZoneIdentity
	}

	return Zone{
		ID:            uuid.New(),
		Identity:      identity,
		Contamination: &IntegerField{},
		Depth:         &IntegerField{},
		Coverage:      &IntegerField{},
	}, nil
}
