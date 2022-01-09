package mongo_model

import "go.mongodb.org/mongo-driver/bson/primitive"

type InputRunway struct {
	Designator    string
	Contamination map[string]int
	Coverage      map[string]int
	Depth         map[string]int
	LooseSand     bool
	LatestVersion int
	MetaData      struct{}
}

type OutputRunway struct {
	ID            primitive.ObjectID `bson:"_id"`
	Designator    string
	Contamination map[string]int
	Coverage      map[string]int
	Depth         map[string]int
	LooseSand     bool
	LatestVersion int
	MetaData      struct{}
}
