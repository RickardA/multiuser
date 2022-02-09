package app

import (
	"github.com/RickardA/multiuser/internal/pkg/domain"
	log "github.com/sirupsen/logrus"
)

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
	versionIsMismatched, err := c.syncHandler.CheckVersionMismatch(input)

	if err != nil {
		return domain.Runway{}, err
	}

	// If version is mismatched, return error
	if versionIsMismatched {
		c.syncHandler.CreateConflict(input)
		return domain.Runway{}, versionMismatchedError
	}

	// If no version mismatch, bump it and then update
	log.WithFields(log.Fields{"id": id}).Info("Bumping version and updating runway")

	updatedRunway, err := c.repository.UpdateRunway(id, input)

	if err != nil {
		return domain.Runway{}, err
	}

	return updatedRunway, nil
}

func (c *Client) DeleteRunwayWithID(id domain.RunwayID) error {
	return nil
}

func (c *Client) GetConflictByID(id domain.ConflictID) (domain.Conflict, error) {
	return domain.Conflict{}, nil
}

func (c *Client) GetConflictForRunway(runwayID domain.RunwayID) (domain.Conflict, error) {
	conflict, err := c.repository.GetConflictForRunway(runwayID)

	if err != nil {
		return domain.Conflict{}, err
	}

	return conflict, nil
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

func (c *Client) ResolveConflict(conflictID domain.ConflictID, resolutionStrategy domain.ResolutionStrategy) (domain.Runway, error) {
	// Get conflict from db
	conflict, err := c.repository.GetConflictByID(conflictID)

	if err != nil {
		return domain.Runway{}, err
	}

	// Get runway from db
	remoteRwy, err := c.repository.GetRunwayByID(conflict.RunwayID)

	if err != nil {
		return domain.Runway{}, err
	}

	//Apply changes
	modifiedRwy, err := c.syncHandler.ApplyChanges(remoteRwy, conflict, resolutionStrategy)

	if err != nil {
		return domain.Runway{}, err
	}

	//Save modified runway and return result
	return c.repository.UpdateRunway(conflict.RunwayID, modifiedRwy)
}
