package memory_runway

import (
	"testing"

	"github.com/RickardA/multiuser/internal/pkg/domain"
	"github.com/RickardA/multiuser/internal/pkg/repository/runway"
)

func TestMemoryRunwayRepository_Add(t *testing.T) {
	repo := New()
	runway, err := domain.CreateRunway("10-23")
	if err != nil {
		t.Error(err)
	}

	repo.Add(runway)
	if len(repo.runways) != 1 {
		t.Errorf("Expected 1 product, got %d", len(repo.runways))
	}
}

func TestMemoryProductRepository_Get(t *testing.T) {
	repo := New()
	existingRunway, err := domain.CreateRunway("10-23")
	if err != nil {
		t.Error(err)
	}

	repo.Add(existingRunway)
	if len(repo.runways) != 1 {
		t.Errorf("Expected 1 runway, got %d", len(repo.runways))
	}

	type testCase struct {
		name        string
		designator  string
		expectedErr error
	}

	testCases := []testCase{
		{
			name:        "Get runway by designator",
			designator:  "10-23",
			expectedErr: nil,
		}, {
			name:        "Get non-existing runway by designator",
			designator:  "10-11",
			expectedErr: runway.ErrRunwayNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := repo.GetByDesignator(tc.designator)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}

		})
	}

}
func TestMemoryRunwayRepository_Update(t *testing.T) {
	repo := New()
	existingRunway, err := domain.CreateRunway("10-23")
	if err != nil {
		t.Error(err)
	}

	repo.Add(existingRunway)
	if len(repo.runways) != 1 {
		t.Errorf("Expected 1 ruwnay, got %d", len(repo.runways))
	}

	existingRunway.LatestVersion = 2

	err = repo.Update(existingRunway)
	if err != nil {
		t.Error(err)
	}

	updatedRunway, err := repo.GetByDesignator("10-23")
	if err != nil {
		t.Error(err)
	}

	if updatedRunway.LatestVersion != 2 {
		t.Errorf("Expected LatestVersion 2, got %d", updatedRunway.LatestVersion)
	}
}
