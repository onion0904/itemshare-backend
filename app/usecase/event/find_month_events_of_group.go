package event

import (
	"context"
	"time"

	eventDomain "github.com/onion0904/CarShareSystem/app/domain/event"
)

type FindMonthEventsOfGroupUseCase struct {
	eventRepo eventDomain.EventRepository
}

func NewFindMonthEventsOfGroupUseCase(
	eventRepo eventDomain.EventRepository,
) *FindMonthEventsOfGroupUseCase {
	return &FindMonthEventsOfGroupUseCase{
		eventRepo: eventRepo,
	}
}

type FindMonthEventsOfGroupUseCaseDto struct {
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

func (uc *FindMonthEventsOfGroupUseCase) Run(ctx context.Context, year int32, month int32, groupID string) ([]*FindMonthEventsOfGroupUseCaseDto, error) {
	events, err := uc.eventRepo.FindMonthEventsOfGroup(ctx, year, month, groupID)
	if err != nil {
		return nil, err
	}
	result := make([]*FindMonthEventsOfGroupUseCaseDto,len(events))
	for _,event := range events{
		result = append(result, &FindMonthEventsOfGroupUseCaseDto{
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
		})
	}
	return result,err
}
