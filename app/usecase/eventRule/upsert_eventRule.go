package eventRule

import (
	"context"

	eventRuleDomain "github.com/onion0904/CarShareSystem/app/domain/eventRule"
)

type UpsertUseCase struct {
	eventRuleRepo eventRuleDomain.EventRuleRepository
}

func NewUpsertUseCase(
	eventRuleRepo eventRuleDomain.EventRuleRepository,
) *UpsertUseCase {
	return &UpsertUseCase{
		eventRuleRepo: eventRuleRepo,
	}
}

type UpsertUseCaseDto struct {
	UserID         string
	ItemID         string
	NormalLimit    int32
	ImportantLimit int32
}

func (uc *UpsertUseCase) Run(ctx context.Context, dto UpsertUseCaseDto) error {
	neventRule, err := eventRuleDomain.NewEventRule(dto.UserID, dto.ItemID, dto.NormalLimit, dto.ImportantLimit)
	if err != nil {
		return err
	}
	err = uc.eventRuleRepo.UpsertEventRule(ctx, neventRule)
	if err != nil {
		return err
	}
	return nil
}
