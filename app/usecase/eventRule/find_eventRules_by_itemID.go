package eventRule

import (
	"context"

	eventRuleDomain "github.com/onion0904/CarShareSystem/app/domain/eventRule"
)

type FindeventRuleByItemIDUseCase struct {
	eventRuleRepo eventRuleDomain.EventRuleRepository
}

func NewFindeventRuleByItemIDUseCase(
	eventRuleRepo eventRuleDomain.EventRuleRepository,
) *FindeventRuleByItemIDUseCase {
	return &FindeventRuleByItemIDUseCase{
		eventRuleRepo: eventRuleRepo,
	}
}

type FindeventRuleByItemIDUseCaseDto struct {
	UserID string
	ItemID string
	NormalLimit int32
	ImportantLimit int32
}

func (uc *FindeventRuleByItemIDUseCase) Run(ctx context.Context, itemID string) (*[]FindeventRuleByItemIDUseCaseDto, error) {
	eventRules, err := uc.eventRuleRepo.FindEventRulesByItemID(ctx, itemID)
	if err != nil {
		return nil, err
	}
	var dto []FindeventRuleByItemIDUseCaseDto
	for _,eventRule := range *eventRules{
		dto = append(dto, 
			FindeventRuleByItemIDUseCaseDto{
				UserID: eventRule.UserID(),
				ItemID: eventRule.ItemID(),
				NormalLimit: eventRule.NormalLimit(),
				ImportantLimit: eventRule.ImportantLimit(),
			})
	}

	return &dto,nil
}
