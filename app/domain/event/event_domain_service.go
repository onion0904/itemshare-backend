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
	if event.date.Before(event.startDate) || event.date.After(event.endDate) {
		return errors.New("イベントが登録可能期間外です")
	}
	oldEvent, err := c.EventRepo.FindDayOfEvent(ctx, event.year, event.month, event.day)

	if err != nil {
		return err
	}

	if oldEvent != nil {
		err = c.validSetableEvents(ctx, oldEvent, event)
		if err != nil {
			return err
		}
	}

	err = c.EventRepo.SaveEvent(ctx, event)
	if err != nil {
		return err
	}

	return nil
}

// eventが被ったとき早い者勝ちにする。eventが被り、片方がimportantをtrueにしている場合はimportantの方を登録。両方importantのときは早い者勝ち。
func (c *EventDomainService) validSetableEvents(ctx context.Context, oldEvent, newEvent *Event) error {
	if oldEvent.important && newEvent.important {
		return errors.New("すでに重要なイベントが登録されています。")
	} else if !oldEvent.important && !newEvent.important {
		return errors.New("すでにイベントが登録されていますが、重要にすれば登録できます。")
	} else if oldEvent.important && !newEvent.important {
		return errors.New("すでに重要なイベントが登録されています。")
	}
	err := c.EventRepo.DeleteEvent(ctx, oldEvent.id)
	if err != nil {
		return err
	}
	return nil
}