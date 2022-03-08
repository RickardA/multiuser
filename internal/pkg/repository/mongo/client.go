package mongo

import (
	"context"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Client struct {
	db  *mongo.Client
	ctx context.Context
}

func NewConnection(ctx context.Context, uri string) (Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return Client{}, err
	}
	err = client.Connect(ctx)
	if err != nil {
		return Client{}, err
	}

	err = client.Ping(ctx, nil)

	if err != nil {
		return Client{}, err
	}

	return Client{
		db:  client,
		ctx: ctx,
	}, nil
}

func (c *Client) Disconnect() error {
	return c.db.Disconnect(c.ctx)
}

func (c *Client) ListDatabaseNames() error {
	databases, err := c.db.ListDatabaseNames(c.ctx, bson.M{})
	if err != nil {
		return err
	}

	log.WithField("Databases", databases).Info("List databases")

	return nil
}
