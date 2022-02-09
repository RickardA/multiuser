package fromgql

import (
	"github.com/RickardA/multiuser/graph/model"
	"github.com/RickardA/multiuser/internal/pkg/domain"
)

func Strategy(input model.Strategy) domain.ResolutionStrategy {
	switch input {
	case model.StrategyApplyLocal:
		return domain.APPLY_LOCAL
	case model.StrategyApplyRemote:
		return domain.APPLY_REMOTE
	default:
		panic("Heeeeeelpppp!")
	}
}
