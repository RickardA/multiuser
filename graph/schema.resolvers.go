package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/RickardA/multiuser/graph/generated"
	"github.com/RickardA/multiuser/graph/model"
	"github.com/RickardA/multiuser/internal/pkg/domain"
	"github.com/RickardA/multiuser/internal/pkg/domain/conv/fromgql"
	"github.com/RickardA/multiuser/internal/pkg/domain/conv/intogql"
	uuid "github.com/nu7hatch/gouuid"
)

func (r *mutationResolver) CreateRunway(ctx context.Context, clientID string, input model.NewRunway) (string, error) {
	rwy := domain.Runway{
		Designator: input.Designator,
	}

	res, err := r.Client.CreateRunway(rwy)

	if err != nil {
		return "", err
	}

	return string(res), nil
}

func (r *mutationResolver) UpdateRunway(ctx context.Context, clientID string, input model.GQRunwayInput) (*model.GQRunway, error) {
	runway := fromgql.Runway(input)
	rwy, err := r.Client.UpdateRunway(domain.RunwayID(input.ID), runway)

	if err != nil {
		return &model.GQRunway{}, err
	}

	return intogql.Runway(rwy), nil
}

func (r *mutationResolver) ResolveConflict(ctx context.Context, clientID string, conflictID string, strategy model.Strategy) (*model.GQRunway, error) {
	rwy, err := r.Client.ResolveConflict(domain.ConflictID(conflictID), fromgql.Strategy(strategy))

	if err != nil {
		return &model.GQRunway{}, err
	}

	return intogql.Runway(rwy), nil
}

func (r *queryResolver) GetRunwayByDesignator(ctx context.Context, clientID string, designator string) (*model.GQRunway, error) {
	rwy, err := r.Client.GetRunwayByDesignator(designator)

	if err != nil {
		return &model.GQRunway{}, err
	}

	return intogql.Runway(rwy), nil
}

func (r *queryResolver) GetRunwayByID(ctx context.Context, clientID string, id string) (*model.GQRunway, error) {
	rwy, err := r.Client.GetRunwayByID(domain.RunwayID(id))

	if err != nil {
		return &model.GQRunway{}, err
	}

	return intogql.Runway(rwy), nil
}

func (r *queryResolver) GetConflictByRunwayID(ctx context.Context, clientID string, id string) (*model.GQConflict, error) {
	conflict, err := r.Client.GetConflictForRunway(domain.RunwayID(id))

	if err != nil {
		return &model.GQConflict{}, err
	}

	return intogql.Conflict(conflict)
}

func (r *queryResolver) SayHello(ctx context.Context) (string, error) {
	id, err := uuid.NewV4()

	if err != nil {
		return "", err
	}

	return id.String(), err
}

func (r *subscriptionResolver) Conflict(ctx context.Context, clientID string, runwayID string) (<-chan *model.GQConflict, error) {
	// Subscription req recieved
	// Check what runway they want to subscribe to
	// Put

	id, err := uuid.NewV4()

	if err != nil {
		return nil, err
	}

	ch := make(chan *model.GQConflict, 1)

	r.Client.Subs[id.String()] = ch
	fmt.Println("Conflict subscription created")

	return ch, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
