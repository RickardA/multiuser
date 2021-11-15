package service

import (
	"fmt"
	"testing"

	"github.com/RickardA/multiuser/pkg/aggregate"
	memory_conflictObj "github.com/RickardA/multiuser/pkg/repository/conflict_obj/memory"
	memory_runway "github.com/RickardA/multiuser/pkg/repository/runway/memory"
)

func TestSyncHandler_CheckVersionMismatch(t *testing.T) {
	repo := memory_runway.New()
	conflictRepo := memory_conflictObj.New()

	syncHandler, err := New(repo, conflictRepo)
	if err != nil {
		t.Error(err)
	}

	localRunway, err := aggregate.CreateRunway("10-23")
	if err != nil {
		t.Error(err)
	}

	remoteRunway, err := aggregate.CreateRunway("10-23")
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
	conflictRepo := memory_conflictObj.New()

	syncHandler, err := New(repo, conflictRepo)
	if err != nil {
		t.Error(err)
	}

	localRunway, err := aggregate.CreateRunway("10-23")
	if err != nil {
		t.Error(err)
	}

	remoteRunway, err := aggregate.CreateRunway("10-23")
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

	conflicts := syncHandler.GetConflictingFields(localRunway)

	fmt.Printf("Conflicting fields: %v\n", conflicts)
}

func TestSyncHandler_ApplyChanges(t *testing.T) {
	repo := memory_runway.New()
	conflictRepo := memory_conflictObj.New()

	syncHandler, err := New(repo, conflictRepo)
	if err != nil {
		t.Error(err)
	}

	localRunway, err := aggregate.CreateRunway("10-23")
	if err != nil {
		t.Error(err)
	}

	remoteRunway, err := aggregate.CreateRunway("10-23")
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

	conflicts := syncHandler.GetConflictingFields(localRunway)

	err = conflictRepo.Add(conflicts)
	if err != nil {
		t.Error(err)
	}

	syncHandler.applyChanges(conflicts.ID, "LOCAL")

	fmt.Printf("Conflicting fields: %v\n", conflicts)
}
