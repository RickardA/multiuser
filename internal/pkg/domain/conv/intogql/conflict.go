package intogql

import (
	"encoding/json"

	"github.com/RickardA/multiuser/graph/model"
	"github.com/RickardA/multiuser/internal/pkg/domain"
)

func Conflict(input domain.Conflict) (*model.GQConflict, error) {
	remoteJson, err := json.Marshal(input.Remote)

	if err != nil {
		return &model.GQConflict{}, err
	}

	localJson, err := json.Marshal(input.Local)

	if err != nil {
		return &model.GQConflict{}, err
	}

	return &model.GQConflict{
		ID:               string(input.ID),
		RunwayID:         string(input.RunwayID),
		ResolutionMethod: input.ResolutionMethod,
		Remote:           string(remoteJson),
		Local:            string(localJson),
	}, nil
}
