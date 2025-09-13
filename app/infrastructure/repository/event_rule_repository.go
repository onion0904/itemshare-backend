package repository

import (
	"context"
	"database/sql"
	"errors"
	"log"

	errDomain "github.com/onion0904/CarShareSystem/app/domain/error"
	"github.com/onion0904/CarShareSystem/app/domain/eventRule"
	"github.com/onion0904/CarShareSystem/app/infrastructure/db"
	dbgen "github.com/onion0904/CarShareSystem/app/infrastructure/db/sqlc/dbgen"
)

type eventRuleRepository struct {
	db *sql.DB
}

func NewEventRuleRepository(db *sql.DB) eventRule.EventRuleRepository {
	return &eventRuleRepository{db: db}
}

func (er *eventRuleRepository) UpsertEventRule(ctx context.Context, eventRule *eventRule.EventRule) error {
	query := db.GetQuery(ctx)

	err := query.UpsertEventRule(ctx, dbgen.UpsertEventRuleParams{
		UserID:         eventRule.UserID(),
		ItemID:         eventRule.ItemID(),
		NormalLimit:    eventRule.NormalLimit(),
		ImportantLimit: eventRule.ImportantLimit(),
	})
	if err != nil {
		return err
	}
	return nil
}

func (er *eventRuleRepository) FindEventRuleByUserAndItem(ctx context.Context, userID, itemID string) (normalLimit, importantLimit int32, err error) {
	query := db.GetQuery(ctx)
	e, err := query.GetEventRuleByUserAndItem(ctx, dbgen.GetEventRuleByUserAndItemParams{
		UserID: userID,
		ItemID: itemID,
	})
	if err != nil {
		log.Printf("Error finding event rule: %v", err) // エラーログ追加
		if errors.Is(err, sql.ErrNoRows) {
			return -1, -1, errDomain.NewError("eventRule not found")
		}
		return -1, -1, err
	}

	log.Printf("Found event rule: normal=%d, important=%d", e.NormalLimit, e.ImportantLimit) // 結果確認用ログ
	return e.NormalLimit, e.ImportantLimit, nil
}

func (er *eventRuleRepository) FindEventRulesByItemID(ctx context.Context, itemID string) (eventRules *[]eventRule.EventRule, err error) {
	query := db.GetQuery(ctx)

	es, err := query.GetEventRulesByItemID(ctx, itemID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errDomain.NewError("eventRule not found")
		}
		return nil, err
	}
	var ers []eventRule.EventRule
	for _, e := range es {
		ne, err := eventRule.Reconstruct(
			e.UserID,
			e.ItemID,
			e.NormalLimit,
			e.ImportantLimit,
		)
		if err != nil {
			return nil, err
		}
		ers = append(ers, *ne)
	}

	return &ers, nil
}
