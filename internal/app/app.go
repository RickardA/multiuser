package app

import (
	"fmt"

	"github.com/RickardA/multiuser/internal/pkg/domain"
	"github.com/RickardA/multiuser/internal/pkg/domain/conv/intogql"
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

func (c *Client) UpdateRunway(id domain.RunwayID, input domain.Runway, clientID string) (domain.Runway, error) {
	versionIsMismatched, err := c.syncHandler.CheckVersionMismatch(input)

	if err != nil {
		return domain.Runway{}, err
	}

	// If version is mismatched, return error
	if versionIsMismatched {
		log.WithFields(log.Fields{"runwayID": id, "clientID": clientID}).Info("Version mismatched, creating conflict and pushing to subscriber")
		conflict, err := c.syncHandler.CreateConflict(input, clientID)

		if err != nil {
			return domain.Runway{}, err
		}

		if c.SendConflictToSubscribers(conflict, id, clientID) != nil {
			return domain.Runway{}, err
		}

		log.Info("Returning in update runway")
		return domain.Runway{}, versionMismatchedError
	}

	// If no version mismatch, bump it and then update
	log.WithFields(log.Fields{"id": id}).Info("Bumping version and updating runway")

	updatedRunway, err := c.repository.UpdateRunway(id, input, clientID)

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

func (c *Client) ResolveConflict(conflictID domain.ConflictID, resolutionStrategy domain.ResolutionStrategy, clientID string) (domain.Runway, error) {
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
	return c.repository.UpdateRunway(conflict.RunwayID, modifiedRwy, clientID)
}

func (c *Client) SendConflictToSubscribers(conflict domain.Conflict, runwayID domain.RunwayID, clientID string) error {
	// Get subscription client and send out conflict
	subID := fmt.Sprintf("%v-%v", clientID, runwayID)
	sub := c.Subs[subID]

	gqlConflict, err := intogql.Conflict(conflict)

	if err != nil {
		log.WithFields(log.Fields{"RunwayID": runwayID, "ClientID": clientID}).Error("Could not send conflict to client on socket")
	}

	sub <- gqlConflict

	return nil
}
