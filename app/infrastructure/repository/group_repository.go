package repository

import (
	"context"
	"database/sql"
	"errors"
	"log"

	errDomain "github.com/onion0904/CarShareSystem/app/domain/error"
	"github.com/onion0904/CarShareSystem/app/domain/group"
	"github.com/onion0904/CarShareSystem/app/infrastructure/db"
	dbgen "github.com/onion0904/CarShareSystem/app/infrastructure/db/sqlc/dbgen"
)

type groupRepository struct {
	db *sql.DB
}

func NewGroupRepository(db *sql.DB) group.GroupRepository {
	return &groupRepository{db: db}
}

func (gr *groupRepository) Update(ctx context.Context, group *group.Group) error {
	query := db.GetQuery(ctx)

	err := query.UpsertGroup(ctx, dbgen.UpsertGroupParams{
		Name: group.Name(),
		ID:   group.ID(),
	})
	if err != nil {
		return err
	}
	return nil
}

func (gr *groupRepository) Save(ctx context.Context, group *group.Group) error {
	query := db.GetQuery(ctx)

	err := query.UpsertGroup(ctx, dbgen.UpsertGroupParams{
		ID:   group.ID(),
		Name: group.Name(),
	})
	if err != nil {
		return err
	}
	return nil
}

func (gr *groupRepository) Delete(ctx context.Context, groupID string) error {
	query := db.GetQuery(ctx)

	err := query.DeleteGroup(ctx, groupID)
	if err != nil {
		return err
	}
	return nil
}

func (gr *groupRepository) FindGroupByID(ctx context.Context, groupID string) (*group.Group, error) {
	DB := db.GetDB()
	tx, err := DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := tx.Rollback(); err != nil && err != sql.ErrTxDone {
			log.Printf("rollback failed: %v", err)
		}
	}()

	query := db.GetQuery(ctx).WithTx(tx)

	g, err := query.FindGroup(ctx, groupID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errDomain.NewError("Group not found")
		}
		return nil, err
	}

	userIDs, err := query.GetUserIDsByGroupID(ctx, groupID)
	if err != nil {
		return nil, err
	}

	eventIDs, err := query.GetEventIDsByGroupID(ctx, groupID)
	if err != nil {
		return nil, err
	}

	ng, err := group.Reconstruct(
		g.ID,
		g.Name,
		userIDs,
		eventIDs,
	)
	if err != nil {
		return nil, err
	}
	ng.SetCreatedAt(g.CreatedAt)
	ng.SetUpdatedAt(g.UpdatedAt)

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return ng, nil
}

func (gr *groupRepository) FindGroupsByUserID(ctx context.Context, userID string) ([]*group.Group, error) {
	DB := db.GetDB()
	tx, err := DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := tx.Rollback(); err != nil && err != sql.ErrTxDone {
			log.Printf("rollback failed: %v", err)
		}
	}()

	query := db.GetQuery(ctx).WithTx(tx)

	groups, err := query.FindGroupsByUserID(ctx,userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errDomain.NewError("Group not found")
		}
		return nil, err
	}

	result := make([]*group.Group,len(groups))

	for _,g := range groups{
		userIDs, err := query.GetUserIDsByGroupID(ctx, g.ID)
		if err != nil {
			return nil, err
		}

		eventIDs, err := query.GetEventIDsByGroupID(ctx, g.ID)
		if err != nil {
			return nil, err
		}

		ng, err := group.Reconstruct(
			g.ID,
			g.Name,
			userIDs,
			eventIDs,
		)
		if err != nil {
			return nil, err
		}
		ng.SetCreatedAt(g.CreatedAt)
		ng.SetUpdatedAt(g.UpdatedAt)

		result = append(result, ng)
	}
	

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return result, nil
}

func (gr *groupRepository) AddUserToGroup(ctx context.Context, groupID string, userID string) error {
	query := db.GetQuery(ctx)

	err := query.AddUserToGroup(ctx, dbgen.AddUserToGroupParams{
		Groupid: groupID,
		Userid:  userID,
	})
	if err != nil {
		return err
	}
	return nil
}

func (gr *groupRepository) AddEventToGroup(ctx context.Context, groupID string, eventID string) error {
	query := db.GetQuery(ctx)

	err := query.AddEventToGroup(ctx, dbgen.AddEventToGroupParams{
		Groupid: groupID,
		Eventid: eventID,
	})
	if err != nil {
		return err
	}
	return nil
}

func (gr *groupRepository) RemoveUserFromGroup(ctx context.Context, groupID string, userID string) error {
	query := db.GetQuery(ctx)

	err := query.RemoveUserFromGroup(ctx, dbgen.RemoveUserFromGroupParams{
		Groupid: groupID,
		Userid:  userID,
	})
	if err != nil {
		return err
	}
	return nil
}
