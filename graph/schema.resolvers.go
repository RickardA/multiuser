package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/RickardA/multiuser/graph/generated"
	"github.com/RickardA/multiuser/graph/model"
	"github.com/RickardA/multiuser/internal/pkg/domain"
	"github.com/RickardA/multiuser/internal/pkg/domain/conv/intogql"
)

func (r *mutationResolver) CreateRunway(ctx context.Context, input model.NewRunway) (string, error) {
	rwy := domain.Runway{
		Designator: input.Designator,
	}

	res, err := r.Client.CreateRunway(rwy)

	if err != nil {
		return "", err
	}

	return string(res), nil
}

func (r *queryResolver) GetRunwayByDesignator(ctx context.Context, designator string) (*model.GQRunway, error) {
	rwy, err := r.Client.GetRunwayByDesignator(designator)

	if err != nil {
		return &model.GQRunway{}, err
	}

	return intogql.Runway(rwy), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
