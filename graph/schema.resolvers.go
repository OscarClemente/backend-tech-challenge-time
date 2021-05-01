package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strconv"

	"github.com/OscarClemente/backend-tech-challenge-time/graph/generated"
	"github.com/OscarClemente/backend-tech-challenge-time/graph/model"
	"github.com/OscarClemente/backend-tech-challenge-time/model/timers"
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
	internalTimers := timers.GetTimers()
	var timersResponse []*model.Timer

	for _, t := range internalTimers {
		timer := model.Timer{
			ID:          strconv.Itoa(t.Id),
			Title:       t.Title,
			TimeElapsed: t.TimeElapsed,
			CreatedAt:   t.CreatedAt,
		}

		timersResponse = append(timersResponse, &timer)
	}
	return timersResponse, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
