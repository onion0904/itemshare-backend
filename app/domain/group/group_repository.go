package group

import "context"

type GroupRepository interface {
	Update(ctx context.Context, group *Group) error
	Save(ctx context.Context, group *Group) error
	Delete(ctx context.Context, groupID string) error
	FindGroupByID(ctx context.Context, groupID string) (group *Group, err error)
	FindGroupsByUserID(ctx context.Context, userID string) (group []*Group, err error)
	AddUserToGroup(ctx context.Context, groupID string, userID string) error
	AddEventToGroup(ctx context.Context, groupID string, eventID string) error
	RemoveUserFromGroup(ctx context.Context, groupID string, userID string) error
}
