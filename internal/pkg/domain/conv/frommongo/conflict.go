package frommongo

import (
	"github.com/RickardA/multiuser/internal/pkg/domain"
	mongo "github.com/RickardA/multiuser/internal/pkg/repository/mongo/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ConflictID(input primitive.ObjectID) domain.ConflictID {
	return domain.ConflictID(input.Hex())
}

func Conflict(input mongo.OutputConflict) (domain.Conflict, error) {
	return domain.Conflict{
		ID:               ConflictID(input.ID),
		RunwayID:         domain.RunwayID(input.RunwayID),
		Remote:           input.Remote,
		Local:            input.Local,
		ResolutionMethod: input.ResolutionMethod,
	}, nil
}
