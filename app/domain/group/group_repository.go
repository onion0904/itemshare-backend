package group

import "context"

type GroupRepository interface {
	Update(ctx context.Context, group *Group) error
	Save(ctx context.Context, group *Group) error
	Delete(ctx context.Context, groupID string) error
	FindGroup(ctx context.Context, groupID string) (group *Group, err error)
	AddUserToGroup(ctx context.Context, groupID string, userID string) error
	AddEventToGroup(ctx context.Context, groupID string, eventID string) error
	RemoveUserFromGroup(ctx context.Context, groupID string, userID string) error
}
