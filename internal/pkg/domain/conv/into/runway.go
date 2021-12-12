package into

import (
	"github.com/RickardA/multiuser/graph/model"
	"github.com/RickardA/multiuser/internal/pkg/domain"
)

func Runway(input domain.Runway) *model.GQRunway {
	version := 0
	return &model.GQRunway{
		Designator:    input.Designator,
		Contamination: mapToGQTuple(input.Contamination),
		Depth:         mapToGQTuple(input.Depth),
		Coverage:      mapToGQTuple(input.Coverage),
		LooseSand:     &input.LooseSand,
		LatestVersion: &version,
	}
}

func mapToGQTuple(val map[string]int) []*model.GQTuple {
	returnArr := []*model.GQTuple{}
	for key, val := range val {
		returnArr = append(returnArr, &model.GQTuple{
			Key:   key,
			Value: val,
		})
	}

	return returnArr
}
