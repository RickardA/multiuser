package domain

import "github.com/google/uuid"

type BooleanField struct {
	ID       uuid.UUID `json:"id"`
	Value    bool      `diff:"value"`
	EditedBy string    `json:"editedBy"`
	Version  int       `diff:"version"`
}

type BooleanFieldComparisionObj struct {
	ID       uuid.UUID `json:"-"`
	Value    bool      `json:"value"`
	EditedBy string    `json:"-"`
	Version  int       `json:"version"`
}
