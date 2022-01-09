package domain

import (
	"errors"
)

var (
	ErrMissingDesignator = errors.New("Missing runway designator")
)

type RunwayID string

type Runway struct {
	ID            RunwayID
	Designator    string
	Contamination map[string]int
	Coverage      map[string]int
	Depth         map[string]int
	LooseSand     bool
	LatestVersion int
	MetaData      struct{}
}

/*
func CreateRunway(designator string) (Runway, error) {
	if designator == "" {
		return Runway{}, ErrMissingDesignator
	}

	var zones map[string]*Zone = make(map[string]*Zone)

	for _, zoneIdentity := range []string{"A", "B", "C"} {
		if zone, err := CreateZone(zoneIdentity); err == nil {
			zones[zoneIdentity] = &zone
		} else {
			return Runway{}, err
		}
	}

	return Runway{
		ID:            primitive.NewObjectID(),
		Designator:    designator,
		Contamination: map[string]int{"A": 0, "B": 0, "C": 0},
		Coverage:      map[string]int{"A": 0, "B": 0, "C": 0},
		Depth:         map[string]int{"A": 0, "B": 0, "C": 0},
		LooseSand:     false,
		LatestVersion: 0,
	}, nil
}
*/
