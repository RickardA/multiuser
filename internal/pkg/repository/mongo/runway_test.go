package mongo

import (
	"reflect"
	"testing"

	"github.com/RickardA/multiuser/internal/pkg/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestClient_CreateRunway(t *testing.T) {
	objID := primitive.NewObjectID()
	tests := []struct {
		name    string
		args    domain.Runway
		want    domain.Runway
		wantErr bool
	}{
		{
			name: "Create runway",
			args: domain.Runway{
				ID:         objID,
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
				ID:         objID,
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
			c, err := NewConnection("mongodb://localhost")

			if err != nil {
				t.Error("Could not create database connection, err %w", err)
			}

			got, err := c.CreateRunway(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.CreateRunway() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			runway, err := c.GetRunwayByID(got)

			if err != nil {
				t.Errorf("Could not get runway by ID, error = %v", err)
			}

			if !reflect.DeepEqual(runway, tt.want) {
				t.Errorf("Client.CreateRunway() = %v, want %v", runway, tt.want)
			}
		})
	}
}
