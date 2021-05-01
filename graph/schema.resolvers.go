package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"
	"strconv"

	"github.com/OscarClemente/backend-tech-challenge-time/graph/generated"
	"github.com/OscarClemente/backend-tech-challenge-time/graph/model"
	"github.com/OscarClemente/backend-tech-challenge-time/model/timers"
)

func timerToGraphql(internalTimer *timers.Timer) *model.Timer {
	timer := &model.Timer{
		ID:          strconv.Itoa(internalTimer.Id),
		Title:       internalTimer.Title,
		TimeElapsed: internalTimer.TimeElapsed,
		CreatedAt:   internalTimer.CreatedAt,
	}

	return timer
}

func (r *mutationResolver) CreateTimer(ctx context.Context, title string) (*model.Timer, error) {
	internalTimer := &timers.Timer{
		Title: title,
		// Other values are initialized by db
	}

	err := internalTimer.Save()
	timer := timerToGraphql(internalTimer)

	return timer, err
}

func (r *mutationResolver) UpdateTimer(ctx context.Context, input model.UpdatedTimer) (*model.Timer, error) {
	internalId, err := strconv.Atoi(input.ID)
	if err != nil {
		panic(err)
	}

	internalTimer := &timers.Timer{
		Id:          internalId,
		Title:       input.Title,
		TimeElapsed: input.TimeElapsed,
	}

	err = internalTimer.Update()
	timer := timerToGraphql(internalTimer)

	return timer, err
}

func (r *mutationResolver) DeleteTimer(ctx context.Context, id string) (bool, error) {
	internalId, err := strconv.Atoi(id)

	if err != nil {
		log.Println(err)
		return false, err
	}

	err = timers.DeleteTimer(internalId)

	return err == nil, err
}

func (r *queryResolver) Timers(ctx context.Context) ([]*model.Timer, error) {
	internalTimers := timers.GetTimers()
	var timersResponse []*model.Timer

	for _, t := range internalTimers {
		timer := timerToGraphql(&t)
		timersResponse = append(timersResponse, timer)
	}

	return timersResponse, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
