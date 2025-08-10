package event

import "context"

type EventRepository interface {
	//EventDomainService経由で使用してください (domain/calendar/calendar_domain_service.go)
	UpsertEvent(ctx context.Context, event *Event) error
	//以下はEventDomainService経由でなくてOKです
	DeleteEvent(ctx context.Context, eventID string) error
	FindEvent(ctx context.Context, eventID string) (*Event, error)
	FindDayEventsOfGroup(ctx context.Context, year, month, day int32, groupID string) ([]*Event, error)
	FindMonthEventsOfGroup(ctx context.Context, year, month int32, groupID string) ([]*Event, error)
	FindWeeklyEvents(ctx context.Context, year, month, day int32, userID string) ([]*Event, error)
}
