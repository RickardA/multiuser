package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/RickardA/multiuser/graph/generated"
	"github.com/RickardA/multiuser/graph/model"
	"github.com/RickardA/multiuser/internal/pkg/domain"
	"github.com/RickardA/multiuser/internal/pkg/domain/conv/into"
)

func (r *mutationResolver) CreateRunway(ctx context.Context, input model.NewRunway) (*model.GQRunway, error) {

	newRunway, err := domain.CreateRunway(input.Designator)

	if err != nil {
		return nil, err
	}

	return into.Runway(newRunway), nil

	// panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetRunwayByDesignator(ctx context.Context, designator string) (*model.GQRunway, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
