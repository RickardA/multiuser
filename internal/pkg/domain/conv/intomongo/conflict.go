package intomongo

import (
	"github.com/RickardA/multiuser/internal/pkg/domain"
	mongo "github.com/RickardA/multiuser/internal/pkg/repository/mongo/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ConflictID(input domain.ConflictID) (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(string(input))
}

func Conflict(input domain.Conflict) (mongo.InputConflict, error) {
	return mongo.InputConflict{
		RunwayID:         string(input.RunwayID),
		Remote:           input.Remote,
		Local:            input.Local,
		ResolutionMethod: input.ResolutionMethod,
	}, nil
}
