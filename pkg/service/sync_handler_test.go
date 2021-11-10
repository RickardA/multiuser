package service

import (
	"testing"

	"github.com/RickardA/multiuser/pkg/aggregate"

	"github.com/RickardA/multiuser/pkg/repository/runway/memory"
)

func TestSyncHandler_CheckVersionMismatch(t *testing.T) {
	repo := memory.New()

	syncHandler, err := New(repo)
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
	repo := memory.New()

	syncHandler, err := New(repo)
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

	syncHandler.GetConflictingFields(localRunway)
	if err != nil {
		t.Error(err)
	}
}
