package mongo

import "fmt"

func NewMongoTestConnection() (Client, func(), error) {
	client, err := NewConnection("mongodb://localhost")
	return client, func() {
		if err := client.Disconnect(); err != nil {
			fmt.Println(fmt.Errorf("could not disconnect database, err %w", err))
		}
	}, err
}
