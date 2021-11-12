package aggregate_test

import (
	"testing"

	"github.com/RickardA/multiuser/pkg/aggregate"
)

func TestRunway_CreateRunway(t *testing.T) {
	type testCase struct {
		test          string
		designator    string
		expectedError error
	}

	testCases := []testCase{
		{
			test:          "should return error if designator is empty",
			designator:    "",
			expectedError: aggregate.ErrMissingDesignator,
		},
		{
			test:          "validvalues",
			designator:    "10-23",
			expectedError: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := aggregate.CreateRunway(tc.designator)
			if err != tc.expectedError {
				t.Errorf("Expected error: %v, got: %v", tc.expectedError, err)
			}
		})
	}

	t.Run("Version Check", func(t *testing.T) {
		rwy, err := aggregate.CreateRunway("10-23")
		if err != nil {
			t.Fatalf("Could not create runway and check version")
		}

		if rwy.LatestVersion != 0 {
			t.Errorf("Expected version to be 0 but got %v", rwy.LatestVersion)
		}
	})

	t.Run("Zone Check", func(t *testing.T) {
		rwy, err := aggregate.CreateRunway("10-23")
		if err != nil {
			t.Fatalf("Could not create runway and check zones")
		}

		if len(rwy.Contamination) != 3 {
			t.Errorf("Expected to find 3 zones, found %v", len(rwy.Contamination))
		}
	})
}
