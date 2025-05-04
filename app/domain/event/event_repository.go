package event

import "context"


type EventRepository interface {
	//EventDomainService経由で使用してください (domain/calendar/calendar_domain_service.go)
	SaveEvent(ctx context.Context, event *Event) error
	//以下はEventDomainService経由でなくてOKです
	DeleteEvent(ctx context.Context , eventID string) error
	FindEvent(ctx context.Context, eventID string) (*Event, error)
	FindMonthEventIDs(ctx context.Context, year int32, month int32) (eventID []string,err error)
}