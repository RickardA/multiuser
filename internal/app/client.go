package app

import (
	"github.com/RickardA/multiuser/graph/model"
	"github.com/RickardA/multiuser/internal/pkg/domain"
)

func NewClient(repo Repository, syncHandler SyncHandler) Client {
	return Client{
		repository:  repo,
		syncHandler: syncHandler,
		Subs:        make(map[string]chan *model.GQConflict),
	}
}

type Client struct {
	repository  Repository
	syncHandler SyncHandler
	Subs        map[string]chan *model.GQConflict
}

type Interface interface {
	Repository
	ResolveConflict(conflictID domain.ConflictID, resolutionStrategy domain.ResolutionStrategy, clientID string) (domain.Runway, error)
}

type Repository interface {
	RunwayRepository
	ConflictRepository
}

type RunwayRepository interface {
	GetRunwayByDesignator(designator string) (domain.Runway, error)
	GetRunwayByID(id domain.RunwayID) (domain.Runway, error)
	CreateRunway(input domain.Runway) (domain.RunwayID, error)
	UpdateRunway(id domain.RunwayID, input domain.Runway, clientID string) (domain.Runway, error)
	DeleteRunwayWithID(id domain.RunwayID) error
}

type ConflictRepository interface {
	GetConflictByID(id domain.ConflictID) (domain.Conflict, error)
	GetConflictForRunway(runwayID domain.RunwayID) (domain.Conflict, error)
	CreateConflict(input domain.Conflict) (domain.ConflictID, error)
	UpdateConflict(input domain.Conflict) (domain.Conflict, error)
	DeleteConflictWithID(id domain.ConflictID) error
}

type SyncHandler interface {
	CheckVersionMismatch(localRunway domain.Runway) (bool, error)
	GetConflictingFields(localRunway domain.Runway, remoteRunway domain.Runway) domain.Conflict
	CreateConflict(localRunway domain.Runway, clientID string) (domain.Conflict, error)
	ApplyChanges(remoteRunway domain.Runway, conflict domain.Conflict, strategy domain.ResolutionStrategy) (domain.Runway, error)
}
