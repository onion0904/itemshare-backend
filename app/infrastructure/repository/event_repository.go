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

func (er *eventRepository) UpsertEvent(ctx context.Context, event *event.Event) error {
	query := db.GetQuery(ctx)

	err := query.UpsertEvent(ctx, dbgen.UpsertEventParams{
		ID:          event.ID(),
		UserID:      event.UserID(),
		ItemID: 	 event.ItemID(),
		Together:    event.Together(),
		Description: event.Description(),
		Year:        event.Year(),
		Month:       event.Month(),
		Day:         event.Day(),
		StartDate:   event.StartDate(),
		EndDate:     event.EndDate(),
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

func (er *eventRepository) FindDayEvents(ctx context.Context, year, month, day int32) ([]*event.Event, error) {
	query := db.GetQuery(ctx)

	InputFindDayOfEventParams := dbgen.FindDayEventsParams{
		Year:  year,
		Month: month,
		Day:   day,
	}
	events, err := query.FindDayEvents(ctx, InputFindDayOfEventParams)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, nil
	}

	result := make([]*event.Event,len(events)) 
	for _,e := range events{
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

func (er *eventRepository) FindDayEventOfGroup(ctx context.Context, year, month, day int32, groupID string) (*event.Event, error){
	query := db.GetQuery(ctx)

	e, err := query.FindDayEvent(ctx,dbgen.FindDayEventParams{
		Year: year,
		Month: month,
		Day: day,
		Groupid: groupID,
	})
	if err != nil {
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

	return ne,err
}
	

func (er *eventRepository) FindMonthEventsOfGroup(ctx context.Context, year, month int32, groupID string) ([]*event.Event, error) {
	query := db.GetQuery(ctx)

	events, err := query.FindMonthEvents(ctx,dbgen.FindMonthEventsParams{
		Year: year,
		Month: month,
		Groupid: groupID,
	})
	if err != nil {
		return nil, err
	}
	result := make([]*event.Event,len(events)) 
	for _,e := range events{
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
