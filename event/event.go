package event

type EventType string

// Event is a domain event marker.
type Event interface {
	EventType() EventType
	SetID(uint64)
	GetID() uint64
	UnmarshalFn() func([]byte) Event
}

// EventMetaData contains common meta data for Events.
type EventMetaData struct {
	ID uint64 `json:"id"`
}

func (md *EventMetaData) SetID(id uint64) {
	md.ID = id
}

func (md *EventMetaData) GetID() uint64 {
	return md.ID
}
