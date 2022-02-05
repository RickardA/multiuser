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
			c, exitFunc, err := NewMongoTestConnection()

			if err != nil {
				t.Error("Could not create database connection, err %w", err)
			}

			defer exitFunc()

			if err := c.ListDatabaseNames(); (err != nil) != tt.wantErr {
				t.Errorf("Client.ListDatabaseNames() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
