package domain

import (
	"github.com/google/uuid"
)

type ConflictObj struct {
	ID               uuid.UUID              `json:"ID"`
	Remote           map[string]interface{} `json:"Remote"`
	Local            map[string]interface{} `json:"Local"`
	ResolutionMethod string                 `json:"ResolutionMethod"`
}
