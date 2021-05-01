package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/OscarClemente/backend-tech-challenge-time/graph/generated"
	"github.com/OscarClemente/backend-tech-challenge-time/graph/model"
)

func (r *mutationResolver) CreateTimer(ctx context.Context, title string) (*model.Timer, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateTimer(ctx context.Context, input model.UpdatedTimer) (*model.Timer, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteTimer(ctx context.Context, id string) (int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Timers(ctx context.Context) ([]*model.Timer, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
