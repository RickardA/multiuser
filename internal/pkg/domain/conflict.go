package domain

type ConflictID string

type Conflict struct {
	ID               ConflictID `json:"ID"`
	RunwayID         RunwayID
	Remote           map[string]interface{}
	Local            map[string]interface{}
	ResolutionMethod string
}
