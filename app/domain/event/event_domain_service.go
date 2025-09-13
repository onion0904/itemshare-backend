package event

import (
	"context"
	"errors"
	"fmt"
	"log"

	domain_eventRule "github.com/onion0904/CarShareSystem/app/domain/eventRule"
)

type EventDomainService struct {
	EventRepo     EventRepository
	EventRuleRepo domain_eventRule.EventRuleRepository
}

func NewEventDomainService(
	EventRepo EventRepository,
	EventRuleRepo domain_eventRule.EventRuleRepository,
) *EventDomainService {
	return &EventDomainService{
		EventRepo:     EventRepo,
		EventRuleRepo: EventRuleRepo,
	}
}

func (c *EventDomainService) SaveEventService(ctx context.Context, event *Event, groupID string) error {
	// イベント期間の制約を確認
	if event.date.Before(event.startDate) || event.date.After(event.endDate) {
		return errors.New("イベントが登録可能期間外です")
	}

	// イベントの予約可能数の確認
	// イベントルールの取得(userIDとitemIDから)
	normal, important, err := c.EventRuleRepo.FindEventRuleByUserAndItem(ctx, event.userID, event.itemID)
	if err != nil {
		log.Printf("Error finding event rule: %v", err) // エラーログ
		return fmt.Errorf("eventRule not found: %w", err)
	}
	log.Println(normal, important)

	var ncount, icount int32

	// 今週に登録してる予約数を確認
	events, err := c.EventRepo.FindWeeklyEvents(ctx, event.year, event.month, event.day, event.userID)
	if err != nil {
		return err
	}
	for _, event := range events {
		if event.important {
			icount++
		} else {
			ncount++
		}
	}
	if event.important {
		if important == icount {
			return errors.New("イベントルールによりこれ以上重要な用事を追加できません。")
		}
	} else {
		if normal == ncount {
			return errors.New("イベントルールによりこれ以上普通な用事を追加できません。")
		}
	}

	// イベントの被りの制約を確認(引数にitemIDを追加)
	oldEvents, _ := c.EventRepo.FindDayEventsOfGroup(ctx, event.year, event.month, event.day, groupID)
	if len(oldEvents) != 0 {
		for _, oe := range oldEvents {
			// Itemが被ってないかの確認
			equal := c.EqualItemEvents(ctx, oe, event)

			if equal {
				// 重要か普通の制約を確認
				err := c.validImportantOrNormal(ctx, oe, event)
				if err != nil {
					return err
				}
			}
		}
	}

	err = c.EventRepo.UpsertEvent(ctx, event)
	if err != nil {
		return err
	}

	return nil
}

func (c *EventDomainService) EqualItemEvents(ctx context.Context, oldEvent, newEvent *Event) bool {
	return oldEvent.itemID == newEvent.itemID
}

func (c *EventDomainService) validImportantOrNormal(ctx context.Context, oldEvent, newEvent *Event) error {
	switch {
	case oldEvent.important && newEvent.important:
		// 両方重要 → エラー
		return errors.New("すでに重要なイベントが登録されています。")

	case !oldEvent.important && !newEvent.important:
		// 両方普通 → エラー
		return errors.New("すでにイベントが登録されていますが、重要にすれば登録できます。")

	case oldEvent.important && !newEvent.important:
		// 既存重要、新規普通 → エラー
		return errors.New("すでに重要なイベントが登録されています。")

	case !oldEvent.important && newEvent.important:
		// 既存普通、新規重要 → 古いイベントを削除して新規を登録
		if err := c.EventRepo.DeleteEvent(ctx, oldEvent.id); err != nil {
			return err
		}
		return nil
	}

	return nil
}
