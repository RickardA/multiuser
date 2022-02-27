package mongo

import (
	"github.com/RickardA/multiuser/internal/pkg/domain"
	"github.com/RickardA/multiuser/internal/pkg/domain/conv/frommongo"
	"github.com/RickardA/multiuser/internal/pkg/domain/conv/intomongo"
	"github.com/RickardA/multiuser/internal/pkg/repository"
	mongo "github.com/RickardA/multiuser/internal/pkg/repository/mongo/model"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var _ repository.RunwayRepository = &Client{}

var RunwayCollectionName = "runways"

func (c *Client) GetRunwayByDesignator(designator string) (domain.Runway, error) {
	coll := c.db.Database("db").Collection(RunwayCollectionName)

	result := coll.FindOne(c.ctx, bson.M{"designator": designator})

	var res mongo.OutputRunway

	bytes, err := result.DecodeBytes()
	if err != nil {
		return domain.Runway{}, err
	}

	err = bson.Unmarshal(bytes, &res)

	if err != nil {
		return domain.Runway{}, err
	}

	return frommongo.Runway(res)
}

func (c *Client) GetRunwayByID(id domain.RunwayID) (domain.Runway, error) {
	coll := c.db.Database("db").Collection(RunwayCollectionName)

	convID, err := intomongo.RunwayID(id)

	if err != nil {
		log.WithField("id", id).Error("Could not convert domain.RunwayID to mongo.RunwayID")
		return domain.Runway{}, err
	}

	result := coll.FindOne(c.ctx, bson.M{"_id": convID})

	var res mongo.OutputRunway

	bytes, err := result.DecodeBytes()
	if err != nil {
		log.WithField("id", id).Error("Could not decode result")
		return domain.Runway{}, err
	}

	err = bson.Unmarshal(bytes, &res)

	if err != nil {
		log.WithField("id", id).Error("Could not unmarshal result")
		return domain.Runway{}, err
	}

	return frommongo.Runway(res)
}

func (c *Client) CreateRunway(input domain.Runway) (domain.RunwayID, error) {
	ipt, err := intomongo.Runway(input)

	if err != nil {
		return "", err
	}

	obj, err := bson.Marshal(ipt)
	if err != nil {
		return "", err
	}

	coll := c.db.Database("db").Collection(RunwayCollectionName)

	result, err := coll.InsertOne(c.ctx, obj)

	if err != nil {
		return "", err
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		if err != nil {
			return "", err
		}

		return frommongo.RunwayID(oid), nil
	}

	return "", repository.ErrCouldNotGetObjectID
}

func (c *Client) UpdateRunway(id domain.RunwayID, input domain.Runway, clientID string) (domain.Runway, error) {
	coll := c.db.Database("db").Collection(RunwayCollectionName)

	convID, err := intomongo.RunwayID(id)

	if err != nil {
		log.WithField("id", id).Error("Could not convert domain.RunwayID to mongo.RunwayID")
		return domain.Runway{}, err
	}

	// Bump latestVersion
	input.LatestVersion += 1

	mongoRunway, err := intomongo.Runway(input)

	if err != nil {
		log.WithField("id", id).Error("Could not convert domain.Runway to mongo.Runway")
		return domain.Runway{}, err
	}

	updateBytes, err := bson.Marshal(mongoRunway)

	if err != nil {
		log.WithField("id", id).Error("Could not convert input to json")
		return domain.Runway{}, err
	}

	var updateJSON bson.M
	err = bson.Unmarshal(updateBytes, &updateJSON)

	if err != nil {
		log.WithField("id", id).Error("Could not unmarshal")
		return domain.Runway{}, err
	}

	result, err := coll.UpdateOne(c.ctx, bson.M{"_id": convID}, bson.M{"$set": updateJSON})

	if err != nil {
		log.WithField("id", id).Error("Could not update runway in db")
		return domain.Runway{}, err
	}

	if result.MatchedCount == 0 {
		log.WithField("id", id).Error("Could not update runway in db, id not found")
		return domain.Runway{}, repository.ErrIDNotFound
	}

	return c.GetRunwayByID(id)
}

func (c *Client) DeleteRunwayWithID(id domain.RunwayID) error {
	coll := c.db.Database("db").Collection(RunwayCollectionName)

	convID, err := intomongo.RunwayID(id)

	if err != nil {
		return err
	}

	_, err = coll.DeleteOne(c.ctx, bson.M{"_id": convID})

	if err != nil {
		return err
	}

	return nil
}
