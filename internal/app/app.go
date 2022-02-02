package app

import "github.com/RickardA/multiuser/internal/pkg/domain"

var _ Interface = &Client{}

func (c *Client) GetRunwayByDesignator(designator string) (domain.Runway, error) {
	rwy, err := c.repository.GetRunwayByDesignator(designator)

	if err != nil {
		return domain.Runway{}, err
	}

	return rwy, nil
}

func (c *Client) GetRunwayByID(id domain.RunwayID) (domain.Runway, error) {
	rwy, err := c.repository.GetRunwayByID(id)

	if err != nil {
		return domain.Runway{}, err
	}

	return rwy, nil
}

func (c *Client) CreateRunway(input domain.Runway) (domain.RunwayID, error) {
	runwayID, err := c.repository.CreateRunway(input)

	if err != nil {
		return domain.RunwayID(""), err
	}

	return runwayID, nil
}

func (c *Client) UpdateRunway(id domain.RunwayID, input domain.Runway) (domain.Runway, error) {
	return domain.Runway{}, nil
}

func (c *Client) DeleteRunwayWithID(id domain.RunwayID) error {
	return nil
}

func (c *Client) GetConflictByID(id domain.ConflictID) (domain.Conflict, error) {
	return domain.Conflict{}, nil
}

func (c *Client) GetConflictForRunway(runwayID domain.RunwayID) (domain.Conflict, error) {
	return domain.Conflict{}, nil
}

func (c *Client) CreateConflict(input domain.Conflict) (domain.ConflictID, error) {
	return domain.ConflictID(""), nil
}

func (c *Client) UpdateConflict(input domain.Conflict) (domain.Conflict, error) {
	return domain.Conflict{}, nil
}

func (c *Client) DeleteConflictWithID(id domain.ConflictID) error {
	return nil
}
