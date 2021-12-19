package mongo

import (
	"github.com/RickardA/multiuser/internal/pkg/domain"
	"github.com/RickardA/multiuser/internal/pkg/repository"
)

var _ repository.RunwayRepository = &Client{}

func (c *Client) GetRunwayByDesignator(designator string) (domain.Runway, error) {
	return domain.Runway{}, repository.ErrNotImplemented
}

func (c *Client) GetRunwayByID(id string) (domain.Runway, error) {
	return domain.Runway{}, repository.ErrNotImplemented
}

func (c *Client) CreateRunway(domain.Runway) (domain.Runway, error) {
	return domain.Runway{}, repository.ErrNotImplemented
}

func (c *Client) UpdateRunway(domain.Runway) (domain.Runway, error) {
	return domain.Runway{}, repository.ErrNotImplemented
}

func (c *Client) DeleteRunwayWithID(id string) error {
	return repository.ErrNotImplemented
}
