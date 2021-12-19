package mongo

import (
	"github.com/RickardA/multiuser/internal/pkg/domain"
	"github.com/RickardA/multiuser/internal/pkg/repository"
)

var _ repository.ConflictRepository = &Client{}

func (c *Client) GetConflictForRunway(designator string) (domain.ConflictObj, error) {
	return domain.ConflictObj{}, repository.ErrNotImplemented
}

func (c *Client) CreateConflictObj(domain.ConflictObj) (domain.ConflictObj, error) {
	return domain.ConflictObj{}, repository.ErrNotImplemented
}

func (c *Client) UpdateConflictObj(domain.ConflictObj) (domain.ConflictObj, error) {
	return domain.ConflictObj{}, repository.ErrNotImplemented
}

func (c *Client) DeleteConflictWithID(id string) error {
	return repository.ErrNotImplemented
}
