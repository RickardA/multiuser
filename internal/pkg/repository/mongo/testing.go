package mongo

import (
	"context"
	"fmt"
)

func NewMongoTestConnection() (Client, func(), error) {
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()
	client, err := NewConnection(ctx, "mongodb://localhost")
	return client, func() {
		if err := client.Disconnect(); err != nil {
			fmt.Println(fmt.Errorf("could not disconnect database, err %w", err))
		}
	}, err
}
