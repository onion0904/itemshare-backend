package eventRule

type EventRule struct {
	userID         string
	itemID         string
	normalLimit    int32
	importantLimit int32
}

func Reconstruct(
	userID string,
	itemID string,
	normalLimit int32,
	importantLimit int32,
) (*EventRule, error) {
	return newEventRule(
		userID,
		itemID,
		normalLimit,
		importantLimit,
	)
}

func NewEventRule(
	userID string,
	itemID string,
	normalLimit int32,
	importantLimit int32,
) (*EventRule, error) {
	return newEventRule(
		userID,
		itemID,
		normalLimit,
		importantLimit,
	)
}

func newEventRule(
	userID string,
	itemID string,
	normalLimit int32,
	importantLimit int32,
) (*EventRule, error) {
	return &EventRule{
		userID:         userID,
		itemID:         itemID,
		normalLimit:    normalLimit,
		importantLimit: importantLimit,
	}, nil
}

func (e *EventRule) UserID() string {
	return e.userID
}

func (e *EventRule) ItemID() string {
	return e.itemID
}

func (e *EventRule) NormalLimit() int32 {
	return e.normalLimit
}

func (e *EventRule) ImportantLimit() int32 {
	return e.importantLimit
}
