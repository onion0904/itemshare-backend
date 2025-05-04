package event

import (
	"context"
	eventDomain "github.com/onion0904/CarShareSystem/app/domain/event"
)

type FindMonthEventUseCase struct {
	eventRepo eventDomain.EventRepository
}

func NewFindMonthEventUseCase(
	eventRepo eventDomain.EventRepository,
) *FindMonthEventUseCase {
	return &FindMonthEventUseCase{
		eventRepo: eventRepo,
	}
}

type FindMonthEventUseCaseDto struct {
	EventIDs []string
}

func (uc *FindMonthEventUseCase) Run(ctx context.Context, year int32, month int32) (*FindMonthEventUseCaseDto, error) {
	eventIDs, err := uc.eventRepo.FindMonthEventIDs(ctx, year, month)
	if err != nil {
		return nil, err
	}
	return &FindMonthEventUseCaseDto{
        EventIDs: eventIDs,
    }, nil
}
