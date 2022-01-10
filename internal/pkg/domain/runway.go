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
