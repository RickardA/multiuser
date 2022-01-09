package mongo

import (
	"reflect"
	"testing"

	"github.com/RickardA/multiuser/internal/pkg/domain"
)

func TestClient_CreateConflict(t *testing.T) {
	tests := []struct {
		name    string
		args    domain.Conflict
		want    domain.Conflict
		wantErr bool
	}{
		{
			name: "Create runway",
			args: domain.Conflict{
				RunwayID:         "test",
				ResolutionMethod: "hejsan",
			},
			want: domain.Conflict{
				RunwayID:         "test",
				ResolutionMethod: "hejsan",
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

			got, err := c.CreateConflict(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.CreateConflict() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			runway, err := c.GetConflictByID(got)

			tt.want.ID = got

			if err != nil {
				t.Errorf("Could not get conflict, error = %v", err)
			}

			if !reflect.DeepEqual(runway, tt.want) {
				t.Errorf("Client.CreateConflict() = %v, want %v", runway, tt.want)
			}
		})
	}
}
