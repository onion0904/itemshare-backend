package repository

import (
	"context"
	"database/sql"
	"errors"

	errDomain "github.com/onion0904/CarShareSystem/app/domain/error"
    "github.com/onion0904/CarShareSystem/app/domain/event"
	"github.com/onion0904/CarShareSystem/app/infrastructure/db"
	dbgen "github.com/onion0904/CarShareSystem/app/infrastructure/db/sqlc/dbgen"
)

type eventRepository struct {
	db *sql.DB
}

func NewEventRepository(db *sql.DB) event.EventRepository {
	return &eventRepository{db: db}
}

func (er *eventRepository)SaveEvent(ctx context.Context, event *event.Event) error {
	query := db.GetQuery(ctx)

	err := query.UpsertEvent(ctx ,dbgen.UpsertEventParams{
		ID:          event.ID(),
		UserID:     event.UserID(),
		Together:    event.Together(),
		Description: event.Description(),
		Year:        event.Year(),
		Month:       event.Month(),
		Day:         event.Day(),
		StartDate:   event.StartDate(),
		EndDate:     event.EndDate(),
	})
	if err!= nil {
		return err
	}
	return nil
}
	
func (er *eventRepository)DeleteEvent(ctx context.Context , eventID string) error {
	query := db.GetQuery(ctx)

	err := query.DeleteEvent(ctx, eventID)
	if err!= nil {
        return err
    }
	return nil
}
	
func (er *eventRepository)FindEvent(ctx context.Context, eventID string) (*event.Event, error) {
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

func (er *eventRepository)FindDayOfEvent(ctx context.Context, year, month,day int32) (*event.Event, error) {
	query := db.GetQuery(ctx)

	InputFindDayOfEventParams := dbgen.FindDayOfEventParams{
		Year: year,
		Month: month,
		Day: day,
	}
	e, err := query.FindDayOfEvent(ctx, InputFindDayOfEventParams)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil,nil
	}

	ne, err := event.Reconstruct(
		e.ID,
        e.UserID,
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
	
func (er *eventRepository)FindMonthEventIDs(ctx context.Context, year int32, month int32) ([]string, error) {
	query := db.GetQuery(ctx)

	eventIDs, err := query.FindMonthEventIDs(ctx, dbgen.FindMonthEventIDsParams{
		Year:  year,
        Month: month,
	})
	if err!= nil {
		return nil, err
	}
	return eventIDs, nil
}