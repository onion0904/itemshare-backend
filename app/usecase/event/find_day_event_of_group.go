package event

import (
	"context"
	"time"

	eventDomain "github.com/onion0904/CarShareSystem/app/domain/event"
)

type FindDayEventsOfGroupUseCase struct {
	eventRepo eventDomain.EventRepository
}

func NewFindDayEventsOfGroupUseCase(
	eventRepo eventDomain.EventRepository,
) *FindDayEventsOfGroupUseCase {
	return &FindDayEventsOfGroupUseCase{
		eventRepo: eventRepo,
	}
}

type FindDayEventsOfGroupUseCaseDto struct {
	ID          string
	UserID      string
	ItemID      string
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

func (uc *FindDayEventsOfGroupUseCase) Run(ctx context.Context, year, month, day int32, groupID string) ([]*FindDayEventsOfGroupUseCaseDto, error) {
	events, err := uc.eventRepo.FindDayEventsOfGroup(ctx, year, month, day, groupID)
	if err != nil {
		return nil, err
	}
	result := make([]*FindDayEventsOfGroupUseCaseDto, 0, len(events))
	for _, event := range events {
		result = append(result, &FindDayEventsOfGroupUseCaseDto{
			ID:          event.ID(),
			UserID:      event.UserID(),
			ItemID:      event.ItemID(),
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
		})
	}

	return result, err
}
