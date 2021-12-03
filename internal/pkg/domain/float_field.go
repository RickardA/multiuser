package domain

import "github.com/google/uuid"

type FloatField struct {
	ID       uuid.UUID `json:"id"`
	Value    float64   `diff:"value"`
	EditedBy string    `json:"editedBy"`
	Version  int       `diff:"version"`
}

type FloatFieldComparisionObj struct {
	ID       uuid.UUID `json:"-"`
	Value    float64   `json:"value"`
	EditedBy string    `json:"-"`
	Version  int       `json:"version"`
}
