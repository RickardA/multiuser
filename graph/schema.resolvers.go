package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/RickardA/multiuser/graph/generated"
	"github.com/RickardA/multiuser/graph/model"
	"github.com/RickardA/multiuser/internal/pkg/domain"
	"github.com/RickardA/multiuser/internal/pkg/domain/conv/fromgql"
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

func (r *mutationResolver) UpdateRunway(ctx context.Context, input model.GQRunwayInput) (*model.GQRunway, error) {
	runway := fromgql.Runway(input)
	rwy, err := r.Client.UpdateRunway(domain.RunwayID(input.ID), runway)

	if err != nil {
		return &model.GQRunway{}, err
	}

	return intogql.Runway(rwy), nil
}

func (r *queryResolver) GetRunwayByDesignator(ctx context.Context, designator string) (*model.GQRunway, error) {
	rwy, err := r.Client.GetRunwayByDesignator(designator)

	if err != nil {
		return &model.GQRunway{}, err
	}

	return intogql.Runway(rwy), nil
}

func (r *queryResolver) GetRunwayByID(ctx context.Context, id string) (*model.GQRunway, error) {
	rwy, err := r.Client.GetRunwayByID(domain.RunwayID(id))

	if err != nil {
		return &model.GQRunway{}, err
	}

	return intogql.Runway(rwy), nil
}

func (r *queryResolver) GetConflictByRunwayID(ctx context.Context, id string) (*model.GQConflict, error) {
	conflict, err := r.Client.GetConflictForRunway(domain.RunwayID(id))

	if err != nil {
		return &model.GQConflict{}, err
	}

	return intogql.Conflict(conflict)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
