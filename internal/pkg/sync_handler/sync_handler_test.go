package sync_handler

import (
	"fmt"
	"testing"

	"github.com/RickardA/multiuser/internal/pkg/domain"
	memory_conflict_repository "github.com/RickardA/multiuser/internal/pkg/repository/conflict/memory"
	memory_runway "github.com/RickardA/multiuser/internal/pkg/repository/runway/memory"
)

func TestSyncHandler_CheckVersionMismatch(t *testing.T) {
	repo := memory_runway.New()
	conflictRepo := memory_conflict_repository.New()

	syncHandler, err := New(repo, conflictRepo)
	if err != nil {
		t.Error(err)
	}

	localRunway, err := domain.CreateRunway("10-23")
	if err != nil {
		t.Error(err)
	}

	remoteRunway, err := domain.CreateRunway("10-23")
	if err != nil {
		t.Error(err)
	}
	err = repo.Add(remoteRunway)
	if err != nil {
		t.Error(err)
	}

	mismatch, err := syncHandler.CheckVersionMismatch(localRunway)
	if err != nil {
		t.Error(err)
	}

	if mismatch {
		t.Errorf("Expected mismatch = false got mismatch = %v", mismatch)
	}

	remoteRunway.LatestVersion = 1

	err = repo.Update(remoteRunway)
	if err != nil {
		t.Error(err)
	}

	mismatch, err = syncHandler.CheckVersionMismatch(localRunway)
	if err != nil {
		t.Error(err)
	}

	if !mismatch {
		t.Errorf("Expected mismatch = true got mismatch = %v", mismatch)
	}
}

func TestSyncHandler_GetConflictingFields(t *testing.T) {
	repo := memory_runway.New()
	conflictRepo := memory_conflict_repository.New()

	syncHandler, err := New(repo, conflictRepo)
	if err != nil {
		t.Error(err)
	}

	localRunway, err := domain.CreateRunway("10-23")
	if err != nil {
		t.Error(err)
	}

	remoteRunway, err := domain.CreateRunway("10-23")
	if err != nil {
		t.Error(err)
	}

	remoteRunway.Depth["A"] = 3
	remoteRunway.LooseSand = true
	remoteRunway.LatestVersion = 2

	err = repo.Add(remoteRunway)
	if err != nil {
		t.Error(err)
	}

	remoteRunway, err = repo.GetByDesignator(localRunway.Designator)

	if err != nil {
		t.Error(err)
	}

	conflicts := syncHandler.GetConflictingFields(localRunway, remoteRunway)

	fmt.Printf("Conflicting fields: %v\n", conflicts)
}

func TestSyncHandler_ApplyChanges(t *testing.T) {
	repo := memory_runway.New()
	conflictRepo := memory_conflict_repository.New()

	syncHandler, err := New(repo, conflictRepo)
	if err != nil {
		t.Error(err)
	}

	localRunway, err := domain.CreateRunway("10-23")
	if err != nil {
		t.Error(err)
	}

	remoteRunway, err := domain.CreateRunway("10-23")
	if err != nil {
		t.Error(err)
	}

	remoteRunway.Depth["A"] = 3
	remoteRunway.LooseSand = true
	remoteRunway.LatestVersion = 2

	fmt.Printf("Runway designator: %v\n", remoteRunway.Designator)
	err = repo.Add(remoteRunway)
	if err != nil {
		t.Error(err)
	}

	remoteRunway, err = repo.GetByDesignator(localRunway.Designator)

	if err != nil {
		t.Error(err)
	}

	conflicts := syncHandler.GetConflictingFields(localRunway, remoteRunway)

	err = conflictRepo.Add(conflicts)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("Conflicting fields: %v\n", conflicts)

	syncHandler.applyChanges(conflicts.ID, "LOCAL")

}
