package fromgql

import (
	"github.com/RickardA/multiuser/graph/model"
	"github.com/RickardA/multiuser/internal/pkg/domain"
)

func Runway(input model.GQRunway) domain.Runway {
	return domain.Runway{
		ID:            domain.RunwayID(input.ID),
		Designator:    input.Designator,
		Contamination: GQTupleToMap(input.Contamination),
		Depth:         GQTupleToMap(input.Depth),
		Coverage:      GQTupleToMap(input.Coverage),
		LooseSand:     *input.LooseSand,
		LatestVersion: 0,
	}
}

func GQTupleToMap(tuples []*model.GQTuple) map[string]int {
	returnMap := make(map[string]int)
	for _, val := range tuples {
		returnMap[val.Key] = val.Value
	}

	return returnMap
}
