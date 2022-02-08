package fromgql

import (
	"github.com/RickardA/multiuser/graph/model"
	"github.com/RickardA/multiuser/internal/pkg/domain"
)

func Runway(input model.GQRunwayInput) domain.Runway {
	return domain.Runway{
		ID:            domain.RunwayID(input.ID),
		Designator:    input.Designator,
		Contamination: GQTupleToMap(input.Contamination),
		Depth:         GQTupleToMap(input.Depth),
		Coverage:      GQTupleToMap(input.Coverage),
		LooseSand:     *input.LooseSand,
		LatestVersion: *input.LatestVersion,
	}
}

func GQTupleToMap(tuples []*model.GQTupleInput) map[string]int {
	returnMap := make(map[string]int)
	for _, val := range tuples {
		returnMap[val.Key] = val.Value
	}

	return returnMap
}
