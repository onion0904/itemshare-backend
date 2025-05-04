package group

type GroupUsecase interface {
	InviteGroup(groupID string) (link string,err error)
}