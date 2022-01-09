package mongo

import (
	"github.com/RickardA/multiuser/internal/pkg/domain"
	"github.com/RickardA/multiuser/internal/pkg/domain/conv/frommongo"
	"github.com/RickardA/multiuser/internal/pkg/domain/conv/intomongo"
	"github.com/RickardA/multiuser/internal/pkg/repository"
	mongo "github.com/RickardA/multiuser/internal/pkg/repository/mongo/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var _ repository.ConflictRepository = &Client{}

var ConflictCollectionName = "conflicts"

func (c *Client) GetConflictByID(id domain.ConflictID) (domain.Conflict, error) {
	coll := c.db.Database("db").Collection(ConflictCollectionName)

	convID, err := intomongo.ConflictID(id)

	if err != nil {
		return domain.Conflict{}, err
	}

	result := coll.FindOne(c.ctx, bson.M{"_id": convID})

	var res mongo.OutputConflict

	bytes, err := result.DecodeBytes()
	if err != nil {
		return domain.Conflict{}, err
	}

	err = bson.Unmarshal(bytes, &res)

	if err != nil {
		return domain.Conflict{}, err
	}

	return frommongo.Conflict(res)
}

func (c *Client) GetConflictForRunway(runwayID domain.RunwayID) (domain.Conflict, error) {
	coll := c.db.Database("db").Collection(ConflictCollectionName)

	result := coll.FindOne(c.ctx, bson.M{"runwayid": runwayID})

	var res mongo.OutputConflict

	bytes, err := result.DecodeBytes()
	if err != nil {
		return domain.Conflict{}, err
	}

	err = bson.Unmarshal(bytes, &res)

	if err != nil {
		return domain.Conflict{}, err
	}

	return frommongo.Conflict(res)
}

func (c *Client) CreateConflict(input domain.Conflict) (domain.ConflictID, error) {
	ipt, err := intomongo.Conflict(input)

	if err != nil {
		return "", err
	}

	obj, err := bson.Marshal(ipt)
	if err != nil {
		return "", err
	}

	coll := c.db.Database("db").Collection(ConflictCollectionName)

	result, err := coll.InsertOne(c.ctx, obj)

	if err != nil {
		return "", err
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		if err != nil {
			return "", err
		}

		return frommongo.ConflictID(oid), nil
	}

	return "", repository.ErrCouldNotGetObjectID
}

func (c *Client) UpdateConflict(domain.Conflict) (domain.Conflict, error) {
	return domain.Conflict{}, repository.ErrNotImplemented
}

func (c *Client) DeleteConflictWithID(id domain.ConflictID) error {
	coll := c.db.Database("db").Collection(ConflictCollectionName)

	convID, err := intomongo.ConflictID(id)

	if err != nil {
		return err
	}

	_, err = coll.DeleteOne(c.ctx, bson.M{"_id": convID})

	if err != nil {
		return err
	}

	return nil
}
