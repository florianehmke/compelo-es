package event

const PlayerCreatedType EventType = "PlayerCreated"

// PlayerCreated event.
type PlayerCreated struct {
	EventMetaData
	GUID        string `json:"guid"`
	Name        string `json:"name"`
	ProjectGUID string `json:"ProjectGuid"`
}

func (e *PlayerCreated) EventType() EventType {
	return PlayerCreatedType
}
