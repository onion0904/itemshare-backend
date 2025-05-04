package event

import (
	"context"
	eventDomain "github.com/onion0904/CarShareSystem/app/domain/event"
)
// eventUsecase 構造体
type SaveEventUsecase struct {
	eventService *eventDomain.EventDomainService
}

// NewCalendarUsecase ファクトリ関数
func NewEventUseCase(
	eventService *eventDomain.EventDomainService,
) *SaveEventUsecase {
	return &SaveEventUsecase{
		eventService: eventService,
	}
}

// AddEventUseCaseDTO ユースケース層で使用する入力データ
type AddEventUseCaseDTO struct {
	UsersID string
	Together bool
	Description string
	Year        int32
	Month       int32
	Day         int32
	Important bool
}


// イベントを追加する
func (uc *SaveEventUsecase) Run(ctx context.Context, dto AddEventUseCaseDTO) (*eventDomain.Event,error) {	
	event, err := eventDomain.NewEvent(dto.UsersID, dto.Together, dto.Description, dto.Year, dto.Month, dto.Day, dto.Important)
	if err != nil {
		return nil,err
	}
	// ドメイン層のサービスを呼び出し
	err = uc.eventService.SaveEventService(ctx, event)
	if err != nil {
		return nil,err
	}
	return uc.eventService.EventRepo.FindEvent(ctx,event.ID())
}
