package aggregate

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrMissingDesignator = errors.New("Missing runway designator")
)

type Runway struct {
	ID            uuid.UUID
	Designator    string
	Contamination map[string]int `json:"contamination"`
	Coverage      map[string]int `json:"coverage"`
	Depth         map[string]int `json:"depth"`
	LooseSand     bool           `json:"looseSand" xml:"sa"`
	LatestVersion int
	MetaData      struct{}
}

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
		ID:            uuid.New(),
		Designator:    designator,
		Contamination: map[string]int{"A": 0, "B": 0, "C": 0},
		Coverage:      map[string]int{"A": 0, "B": 0, "C": 0},
		Depth:         map[string]int{"A": 0, "B": 0, "C": 0},
		LooseSand:     false,
		LatestVersion: 0,
	}, nil
}
