package intomongo

import (
	"github.com/RickardA/multiuser/internal/pkg/domain"
	mongo "github.com/RickardA/multiuser/internal/pkg/repository/mongo/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RunwayID(input domain.RunwayID) (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(string(input))
}

func Runway(input domain.Runway) (mongo.InputRunway, error) {
	return mongo.InputRunway{
		Designator:    input.Designator,
		Contamination: input.Contamination,
		Depth:         input.Depth,
		Coverage:      input.Coverage,
		LooseSand:     input.LooseSand,
		LatestVersion: input.LatestVersion,
	}, nil
}
