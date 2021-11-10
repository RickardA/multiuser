package entity

import "github.com/google/uuid"

type IntegerField struct {
	ID       uuid.UUID `json:"id"`
	Value    int       `json:"value"`
	EditedBy string    `json:"editedBy"`
	Version  int       `json:"version"`
}

type IntegerFieldComparisionObj struct {
	ID       uuid.UUID `json:"-"`
	Value    int       `json:"value"`
	EditedBy string    `json:"-"`
	Version  int       `json:"version"`
}
