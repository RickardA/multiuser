package service

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/RickardA/multiuser/pkg/aggregate"
	"github.com/RickardA/multiuser/pkg/repository/runway"
)

type SyncHandlerService interface {
	New(db runway.RunwayRepository) (SyncHandler, error)
	CheckVersionMismatch(localRunway aggregate.Runway) (bool, error)
	GetConflictingFields(localRunway aggregate.Runway)
}

type SyncHandler struct {
	db runway.RunwayRepository
}

func New(db runway.RunwayRepository) (SyncHandler, error) {
	return SyncHandler{
		db: db,
	}, nil
}

func (s SyncHandler) CheckVersionMismatch(localRunway aggregate.Runway) (bool, error) {
	remoteRunway, err := s.db.GetByDesignator(localRunway.Designator)

	if err != nil {
		return false, err
	}

	if remoteRunway.LatestVersion == localRunway.LatestVersion {
		return false, nil
	}

	return true, nil
}

func (s SyncHandler) GetConflictingFields(localRunway aggregate.Runway) {
	remoteRunway, err := s.db.GetByDesignator(localRunway.Designator)

	if err != nil {
		os.Exit(1)
	}

	localJSONRunway, err := json.MarshalIndent(remoteRunway, "", "    ")

	if err != nil {
		os.Exit(1)
	}

	fmt.Println(string(localJSONRunway))

	var runwayCompObj aggregate.RunwayComparisionObj
	err = json.Unmarshal(localJSONRunway, &runwayCompObj)

	if err != nil {
		os.Exit(1)
	}

	fmt.Println(runwayCompObj)

	t, err := json.MarshalIndent(runwayCompObj, "", "    ")

	if err != nil {
		os.Exit(1)
	}

	fmt.Println(string(t))
}
