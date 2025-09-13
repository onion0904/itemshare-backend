package eventRule

import "context"

type EventRuleRepository interface {
	// Itemが追加されたときとUserがグループに追加されたときとRule変更時に実行
	UpsertEventRule(ctx context.Context, eventRule *EventRule) error
	FindEventRuleByUserAndItem(ctx context.Context, userID, itemID string) (normalLimit, importantLimit int32, err error)
	FindEventRulesByItemID(ctx context.Context, itemID string) (eventRules *[]EventRule, err error)
}
