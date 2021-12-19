package mongo

import (
	"testing"
)

func TestClient_ListDatabaseNames(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "List Databasenames",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, err := NewConnection("mongodb://localhost")

			if err != nil {
				t.Errorf("Could not create new connection to db, error %w", err)
			}

			if err := c.ListDatabaseNames(); (err != nil) != tt.wantErr {
				t.Errorf("Client.ListDatabaseNames() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
