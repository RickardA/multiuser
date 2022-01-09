package mongo_model

import "go.mongodb.org/mongo-driver/bson/primitive"

type InputConflict struct {
	RunwayID         string
	Remote           map[string]interface{}
	Local            map[string]interface{}
	ResolutionMethod string
}

type OutputConflict struct {
	ID               primitive.ObjectID `bson:"_id"`
	RunwayID         string
	Remote           map[string]interface{}
	Local            map[string]interface{}
	ResolutionMethod string
}
