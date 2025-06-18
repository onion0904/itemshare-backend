package event

import (
	"context"
	"time"

	eventDomain "github.com/onion0904/CarShareSystem/app/domain/event"
)

type FindEventUseCase struct {
	eventRepo eventDomain.EventRepository
}

func NewFindEventUseCase(
	eventRepo eventDomain.EventRepository,
) *FindEventUseCase {
	return &FindEventUseCase{
		eventRepo: eventRepo,
	}
}

type FindEventUseCaseDto struct {
	ID          string
	UserID      string
	Together    bool
	Description string
	Year        int32
	Month       int32
	Day         int32
	Date        time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	StartDate   time.Time
	EndDate     time.Time
	Important   bool
}

func (uc *FindEventUseCase) Run(ctx context.Context, eventID string) (*FindEventUseCaseDto, error) {
	event, err := uc.eventRepo.FindEvent(ctx, eventID)
	if err != nil {
		return nil, err
	}
	return &FindEventUseCaseDto{
		ID:          event.ID(),
		UserID:      event.UserID(),
		Together:    event.Together(),
		Description: event.Description(),
		Year:        event.Year(),
		Month:       event.Month(),
		Day:         event.Day(),
		Date:        event.Date(),
		CreatedAt:   event.CreatedAt(),
		UpdatedAt:   event.UpdatedAt(),
		StartDate:   event.StartDate(),
		EndDate:     event.EndDate(),
		Important:   event.Important(),
	}, nil
}
