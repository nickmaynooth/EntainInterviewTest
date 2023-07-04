package service

import (
	"git.neds.sh/matty/entain/sports/db"
	"git.neds.sh/matty/entain/sports/proto/sports"
	"golang.org/x/net/context"
)

type Racing interface {
	// ListEvents will return a collection of events.
	ListEvents(ctx context.Context, in *sports.ListEventsRequest) (*sports.ListEventsResponse, error)

	// FetchEventById will return a race with the matching Id.
	FetchEventById(ctx context.Context, in *sports.EventRequestById) (*sports.Event, error)
}

// sportsService implements the Racing interface.
type sportsService struct {
	eventsRepo db.EventsRepo
}

// NewRacingService instantiates and returns a new sportsService.
func NewRacingService(eventsRepo db.EventsRepo) Racing {
	return &sportsService{eventsRepo}
}

func (s *sportsService) ListEvents(ctx context.Context, in *sports.ListEventsRequest) (*sports.ListEventsResponse, error) {
	events, err := s.eventsRepo.List(in.Filter)
	if err != nil {
		return nil, err
	}

	return &sports.ListEventsResponse{Events: events}, nil
}

func (s *sportsService) FetchEventById(ctx context.Context, in *sports.EventRequestById) (*sports.Event, error) {
	id := in.Id
	race, err := s.eventsRepo.FetchById(&id)
	if err != nil {
		return nil, err
	}

	return race, nil
}
