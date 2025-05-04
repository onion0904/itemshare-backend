package event

import (
	"context"
	"errors"
)

type EventDomainService struct {
	EventRepo EventRepository
}

func NewEventDomainService(
	EventRepo EventRepository,
) *EventDomainService {
	return &EventDomainService{
		EventRepo: EventRepo,
	}
}

func (c *EventDomainService) SaveEventService(ctx context.Context, event *Event) error {
	// イベント期間の制約を確認
	if event.date.Before(event.startDate) || event.date.After(event.endDate){
		return errors.New("イベントが登録可能期間外です")
	}

	// 現在のイベント数を取得
	events, err := c.EventRepo.FindMonthEventIDs(ctx, event.year, event.month)
	if err != nil {
		return err
	}

	// イベント数の制約を確認
	if !c.validNumEvents(ctx,events) {
		return errors.New("イベントの最大数を超えています")
	}

	err = c.EventRepo.SaveEvent(ctx, event)
	if err != nil {
		return err
	}

	return nil
}

func (c *EventDomainService) validNumEvents (ctx context.Context , eventIDs []string) bool {
	var importantEvent int
	var nimportantEvent int
	for _, eventID := range eventIDs {
		event, err := c.EventRepo.FindEvent(ctx, eventID)
		if err!= nil {
            return false
        }
		if event.important{
			importantEvent++
		}else{
			nimportantEvent++
		}
	}
	if importantEvent < MaxImportantEvents && (nimportantEvent+importantEvent) < MaxEvents{
		return true
	}else if importantEvent >= MaxImportantEvents && (nimportantEvent+importantEvent) >= MaxEvents{
		return false
	}else {
		return false
	}
}

const (
	MaxEvents int = 4
	MaxImportantEvents int = 2
)