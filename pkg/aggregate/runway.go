package aggregate

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrMissingDesignator = errors.New("Missing runway designator")
)

type Runway struct {
	ID            uuid.UUID        `json:"id"`
	Designator    string           `json:"designator"`
	Zones         map[string]*Zone `json:"zones"`
	LatestVersion int              `json:"latestVersion"`
}

type RunwayComparisionObj struct {
	ID            uuid.UUID                      `json:"-"`
	Designator    string                         `json:"-"`
	Zones         map[string]*ZoneComparisionObj `json:"zones"`
	LatestVersion int                            `json:"latestVersion"`
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
		Zones:         zones,
		LatestVersion: 0,
	}, nil
}
