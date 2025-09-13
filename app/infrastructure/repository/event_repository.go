package repository

import (
	"context"
	"database/sql"
	"errors"

	errDomain "github.com/onion0904/CarShareSystem/app/domain/error"
	"github.com/onion0904/CarShareSystem/app/domain/event"
	"github.com/onion0904/CarShareSystem/app/infrastructure/db"
	dbgen "github.com/onion0904/CarShareSystem/app/infrastructure/db/sqlc/dbgen"
	pkgTime "github.com/onion0904/CarShareSystem/pkg/time"
)

type eventRepository struct {
	db *sql.DB
}

func NewEventRepository(db *sql.DB) event.EventRepository {
	return &eventRepository{db: db}
}

func (er *eventRepository) UpsertEvent(ctx context.Context, event *event.Event) error {
	query := db.GetQuery(ctx)

	err := query.UpsertEvent(ctx, dbgen.UpsertEventParams{
		ID:          event.ID(),
		UserID:      event.UserID(),
		ItemID:      event.ItemID(),
		Together:    event.Together(),
		Description: event.Description(),
		Year:        event.Year(),
		Month:       event.Month(),
		Day:         event.Day(),
		Date:        event.Date(),
		StartDate:   event.StartDate(),
		EndDate:     event.EndDate(),
		Important:   event.Important(),
	})
	if err != nil {
		return err
	}
	return nil
}

func (er *eventRepository) DeleteEvent(ctx context.Context, eventID string) error {
	query := db.GetQuery(ctx)

	err := query.DeleteEvent(ctx, eventID)
	if err != nil {
		return err
	}
	return nil
}

func (er *eventRepository) FindEvent(ctx context.Context, eventID string) (*event.Event, error) {
	query := db.GetQuery(ctx)

	e, err := query.FindEvent(ctx, eventID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errDomain.NewError("Event not found")
		}
		return nil, err
	}

	ne, err := event.Reconstruct(
		e.ID,
		e.UserID,
		e.ItemID,
		e.Together,
		e.Description,
		e.Year,
		e.Month,
		e.Day,
		e.Date,
		e.StartDate,
		e.EndDate,
		e.Important,
	)
	if err != nil {
		return nil, err
	}
	ne.SetCreatedAt(e.CreatedAt)
	ne.SetUpdatedAt(e.UpdatedAt)
	return ne, nil
}

func (er *eventRepository) FindDayEventsOfGroup(ctx context.Context, year, month, day int32, groupID string) ([]*event.Event, error) {
	query := db.GetQuery(ctx)

	events, err := query.FindDayEvents(ctx, dbgen.FindDayEventsParams{
		Year:    year,
		Month:   month,
		Day:     day,
		Groupid: groupID,
	})
	if err != nil {
		return nil, err
	}
	result := make([]*event.Event, 0, len(events))
	for _, e := range events {
		ne, err := event.Reconstruct(
			e.ID,
			e.UserID,
			e.ItemID,
			e.Together,
			e.Description,
			e.Year,
			e.Month,
			e.Day,
			e.Date,
			e.StartDate,
			e.EndDate,
			e.Important,
		)

		if err != nil {
			return nil, err
		}
		ne.SetCreatedAt(e.CreatedAt)
		ne.SetUpdatedAt(e.UpdatedAt)

		result = append(result, ne)
	}

	return result, err
}

func (er *eventRepository) FindMonthEventsOfGroup(ctx context.Context, year, month int32, groupID string) ([]*event.Event, error) {
	query := db.GetQuery(ctx)

	events, err := query.FindMonthEvents(ctx, dbgen.FindMonthEventsParams{
		Year:    year,
		Month:   month,
		Groupid: groupID,
	})
	if err != nil {
		return nil, err
	}
	result := make([]*event.Event, 0, len(events))
	for _, e := range events {
		ne, err := event.Reconstruct(
			e.ID,
			e.UserID,
			e.ItemID,
			e.Together,
			e.Description,
			e.Year,
			e.Month,
			e.Day,
			e.Date,
			e.StartDate,
			e.EndDate,
			e.Important,
		)

		if err != nil {
			return nil, err
		}
		ne.SetCreatedAt(e.CreatedAt)
		ne.SetUpdatedAt(e.UpdatedAt)

		result = append(result, ne)
	}
	return result, nil
}

func (er *eventRepository) FindWeeklyEvents(ctx context.Context, year, month, day int32, userID string) ([]*event.Event, error) {
	query := db.GetQuery(ctx)

	events, err := query.FindWeeklyEvents(ctx, dbgen.FindWeeklyEventsParams{
		Userid:     userID,
		Searchdate: pkgTime.CreateEventDate(year, month, day),
	})
	if err != nil {
		return nil, err
	}
	result := make([]*event.Event, 0, len(events))
	for _, e := range events {
		ne, err := event.Reconstruct(
			e.ID,
			e.UserID,
			e.ItemID,
			e.Together,
			e.Description,
			e.Year,
			e.Month,
			e.Day,
			e.Date,
			e.StartDate,
			e.EndDate,
			e.Important,
		)

		if err != nil {
			return nil, err
		}
		ne.SetCreatedAt(e.CreatedAt)
		ne.SetUpdatedAt(e.UpdatedAt)

		result = append(result, ne)
	}
	return result, nil
}
