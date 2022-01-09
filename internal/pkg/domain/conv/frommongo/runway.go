package frommongo

import (
	"github.com/RickardA/multiuser/internal/pkg/domain"
	mongo "github.com/RickardA/multiuser/internal/pkg/repository/mongo/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RunwayID(input primitive.ObjectID) domain.RunwayID {
	return domain.RunwayID(input.Hex())
}

func Runway(input mongo.OutputRunway) (domain.Runway, error) {
	return domain.Runway{
		ID:            RunwayID(input.ID),
		Designator:    input.Designator,
		Contamination: input.Contamination,
		Depth:         input.Depth,
		Coverage:      input.Coverage,
		LooseSand:     input.LooseSand,
		LatestVersion: input.LatestVersion,
	}, nil
}
