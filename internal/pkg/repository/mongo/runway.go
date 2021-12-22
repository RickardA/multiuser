package mongo

import (
	"fmt"

	"github.com/RickardA/multiuser/internal/pkg/domain"
	"github.com/RickardA/multiuser/internal/pkg/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var _ repository.RunwayRepository = &Client{}

var CollectionName = "runways"

func (c *Client) GetRunwayByDesignator(designator string) (domain.Runway, error) {
	return domain.Runway{}, repository.ErrNotImplemented
}

func (c *Client) GetRunwayByID(id primitive.ObjectID) (domain.Runway, error) {
	fmt.Println("asdas")
	fmt.Println(id.String())
	coll := c.db.Database("db").Collection(CollectionName)

	result := coll.FindOne(c.ctx, bson.M{"_id": id})

	var res domain.Runway

	bytes, err := result.DecodeBytes()
	if err != nil {
		return domain.Runway{}, err
	}

	err = bson.Unmarshal(bytes, &res)

	if err != nil {
		return domain.Runway{}, err
	}

	return res, nil
}

func (c *Client) CreateRunway(input domain.Runway) (primitive.ObjectID, error) {
	obj, err := bson.Marshal(input)
	if err != nil {
		return primitive.ObjectID{}, err
	}

	coll := c.db.Database("db").Collection(CollectionName)

	result, err := coll.InsertOne(c.ctx, obj)

	if err != nil {
		return primitive.ObjectID{}, err
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		fmt.Println("asdas")
		fmt.Println(oid.String())

		if err != nil {
			return primitive.ObjectID{}, err
		}

		return oid, nil
	}

	return primitive.ObjectID{}, repository.ErrCouldNotGetObjectID
}

func (c *Client) UpdateRunway(input domain.Runway) (domain.Runway, error) {
	return domain.Runway{}, repository.ErrNotImplemented
}

func (c *Client) DeleteRunwayWithID(id string) error {
	return repository.ErrNotImplemented
}
