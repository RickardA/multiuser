package mongo

import (
	"reflect"
	"testing"

	"github.com/RickardA/multiuser/internal/pkg/domain"
)

func TestClient_CreateRunway(t *testing.T) {
	tests := []struct {
		name    string
		args    domain.Runway
		want    domain.Runway
		wantErr bool
	}{
		{
			name: "Create runway",
			args: domain.Runway{
				Designator: "12-30",
				Contamination: map[string]int{
					"a": 0,
					"b": 12,
					"c": 12,
				},
				Coverage: map[string]int{
					"a": 0,
					"b": 12,
					"c": 12,
				},
				Depth: map[string]int{
					"a": 0,
					"b": 12,
					"c": 12,
				},
				LatestVersion: 12,
			},
			want: domain.Runway{
				Designator: "12-30",
				Contamination: map[string]int{
					"a": 0,
					"b": 12,
					"c": 12,
				},
				Coverage: map[string]int{
					"a": 0,
					"b": 12,
					"c": 12,
				},
				Depth: map[string]int{
					"a": 0,
					"b": 12,
					"c": 12,
				},
				LatestVersion: 12,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, exitFunc, err := NewMongoTestConnection()

			if err != nil {
				t.Error("Could not create database connection, err %w", err)
			}

			defer exitFunc()

			tt.want.ID, err = c.CreateRunway(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.CreateRunway() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			runway, err := c.GetRunwayByID(tt.want.ID)

			if err != nil {
				t.Errorf("Could not get runway by ID, error = %v", err)
			}

			if !reflect.DeepEqual(runway, tt.want) {
				t.Errorf("Client.CreateRunway() = %v, want %v", runway, tt.want)
			}
		})
	}
}

func TestClient_DeleteRunwayWithID(t *testing.T) {
	tests := []struct {
		name    string
		args    domain.Runway
		delete  domain.RunwayID
		want    error
		wantErr bool
	}{
		{
			name: "Delete runway",
			args: domain.Runway{
				Designator: "12-30",
				Contamination: map[string]int{
					"a": 0,
					"b": 12,
					"c": 12,
				},
				Coverage: map[string]int{
					"a": 0,
					"b": 12,
					"c": 12,
				},
				Depth: map[string]int{
					"a": 0,
					"b": 12,
					"c": 12,
				},
				LatestVersion: 12,
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, exitFunc, err := NewMongoTestConnection()

			if err != nil {
				t.Error("Could not create database connection, err %w", err)
			}

			defer exitFunc()

			deleteID, err := c.CreateRunway(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.CreateRunway() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err := c.DeleteRunwayWithID(deleteID); (err != nil) != tt.wantErr {
				t.Errorf("Client.DeleteRunwayWithID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_UpdateRunway(t *testing.T) {
	tests := []struct {
		name    string
		input   domain.Runway
		update  domain.Runway
		wantErr bool
	}{
		{
			name: "Update runway",
			input: domain.Runway{
				Designator: "12-30",
				Contamination: map[string]int{
					"a": 0,
					"b": 12,
					"c": 12,
				},
				Coverage: map[string]int{
					"a": 0,
					"b": 12,
					"c": 12,
				},
				Depth: map[string]int{
					"a": 0,
					"b": 12,
					"c": 12,
				},
				LatestVersion: 12,
			},
			update: domain.Runway{
				Designator: "12-30",
				Contamination: map[string]int{
					"a": 0,
					"b": 1,
					"c": 1,
				},
				Coverage: map[string]int{
					"a": 0,
					"b": 122,
					"c": 122,
				},
				Depth: map[string]int{
					"a": 0,
					"b": 1234,
					"c": 1234,
				},
				LatestVersion: 13,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, exitFunc, err := NewMongoTestConnection()

			if err != nil {
				t.Error("Could not create database connection, err %w", err)
			}

			defer exitFunc()

			got, err := c.CreateRunway(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.CreateRunway() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			runway, err := c.UpdateRunway(got, tt.update)

			runway.ID = ""

			if err != nil {
				t.Errorf("Could not update runway by ID, error = %v", err)
			}

			if !reflect.DeepEqual(runway, tt.update) {
				t.Errorf("Client.UpdateRunway() = %v, want %v", runway, tt.update)
			}
		})
	}
}

func TestClient_GetRunwayByDesignator(t *testing.T) {
	tests := []struct {
		name    string
		args    domain.Runway
		want    domain.Runway
		wantErr bool
	}{
		{
			name: "Get runway by designator",
			args: domain.Runway{
				Designator: "12-30",
				Contamination: map[string]int{
					"a": 0,
					"b": 12,
					"c": 12,
				},
				Coverage: map[string]int{
					"a": 0,
					"b": 12,
					"c": 12,
				},
				Depth: map[string]int{
					"a": 0,
					"b": 12,
					"c": 12,
				},
				LatestVersion: 12,
			},
			want: domain.Runway{
				Designator: "12-30",
				Contamination: map[string]int{
					"a": 0,
					"b": 12,
					"c": 12,
				},
				Coverage: map[string]int{
					"a": 0,
					"b": 12,
					"c": 12,
				},
				Depth: map[string]int{
					"a": 0,
					"b": 12,
					"c": 12,
				},
				LatestVersion: 12,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, exitFunc, err := NewMongoTestConnection()

			if err != nil {
				t.Error("Could not create database connection, err %w", err)
			}

			defer exitFunc()

			_, err = c.CreateRunway(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.CreateRunway() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			runway, err := c.GetRunwayByDesignator(tt.args.Designator)

			runway.ID = ""

			if err != nil {
				t.Errorf("Could not get runway by Designator, error = %v", err)
			}

			if !reflect.DeepEqual(runway, tt.want) {
				t.Errorf("Client.GetRunwayByDesignator() = %v, want %v", runway, tt.want)
			}
		})
	}
}
