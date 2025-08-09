package event

import (
	"context"
	"time"

	eventDomain "github.com/onion0904/CarShareSystem/app/domain/event"
)

type FindDayEventOfGroupUseCase struct {
	eventRepo eventDomain.EventRepository
}

func NewFindDayEventOfGroupUseCase(
	eventRepo eventDomain.EventRepository,
) *FindDayEventOfGroupUseCase {
	return &FindDayEventOfGroupUseCase{
		eventRepo: eventRepo,
	}
}

type FindDayEventOfGroupUseCaseDto struct {
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

func (uc *FindDayEventOfGroupUseCase) Run(ctx context.Context, year, month, day int32, groupID string) (*FindDayEventOfGroupUseCaseDto, error) {
	e, err := uc.eventRepo.FindDayEventOfGroup(ctx,year,month,day,groupID)
	if err != nil {
		return nil, err
	}
	result := &FindDayEventOfGroupUseCaseDto{
		ID:          e.ID(),
		UserID:      e.UserID(),
		Together:    e.Together(),
		Description: e.Description(),
		Year:        e.Year(),
		Month:       e.Month(),
		Day:         e.Day(),
		Date:        e.Date(),
		CreatedAt:   e.CreatedAt(),
		UpdatedAt:   e.UpdatedAt(),
		StartDate:   e.StartDate(),
		EndDate:     e.EndDate(),
		Important:   e.Important(),
	}
	return result,err
}
