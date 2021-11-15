// Package memory is a in memory implementation of the ProductRepository interface.
package memory_conflictObj

import (
	"testing"

	"github.com/RickardA/multiuser/pkg/aggregate"
	conflictObj "github.com/RickardA/multiuser/pkg/repository/conflict_obj"
	"github.com/google/uuid"
)

func createConflictExample() aggregate.ConflictObj {
	return aggregate.ConflictObj{
		ID: uuid.New(),
		Remote: map[string]interface{}{
			"LatestVersion": 0,
			"LooseSand":     false,
			"Depth":         map[string]int{"A": 2},
		},
		Local: map[string]interface{}{
			"LatestVersion": 2,
			"LooseSand":     true,
			"Depth":         map[string]int{"A": 0},
		},
		ResolutionMethod: "LOCAL",
	}
}

func TestMemoryRunwayRepository_Add(t *testing.T) {
	repo := New()

	conflict := createConflictExample()

	repo.Add(conflict)
	if len(repo.conflicts) != 1 {
		t.Errorf("Expected 1 product, got %d", len(repo.conflicts))
	}
}

func TestMemoryProductRepository_Get(t *testing.T) {
	repo := New()
	conflict := createConflictExample()

	repo.Add(conflict)
	if len(repo.conflicts) != 1 {
		t.Errorf("Expected 1 conflictOBJ, got %d", len(repo.conflicts))
	}

	type testCase struct {
		name        string
		id          uuid.UUID
		expectedErr error
	}

	testCases := []testCase{
		{
			name:        "Get conflict by ID",
			id:          conflict.ID,
			expectedErr: nil,
		}, {
			name:        "Get non-existing conflict by ID",
			id:          uuid.New(),
			expectedErr: conflictObj.ErrConflictNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := repo.GetByID(tc.id)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}

		})
	}

}

func TestMemoryRunwayRepository_Delete(t *testing.T) {
	repo := New()
	conflict := createConflictExample()

	repo.Add(conflict)
	if len(repo.conflicts) != 1 {
		t.Errorf("Expected 1 conflictOBJ, got %d", len(repo.conflicts))
	}

	type testCase struct {
		name        string
		id          uuid.UUID
		expectedErr error
	}

	testCases := []testCase{
		{
			name:        "Get deleted conflict by ID",
			id:          conflict.ID,
			expectedErr: conflictObj.ErrConflictNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := repo.Delete(tc.id)
			_, err = repo.GetByID(tc.id)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}

		})
	}
}
