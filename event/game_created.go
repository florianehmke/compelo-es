package event

const GameCreatedType EventType = "GameCreated"

// GameCreated event.
type GameCreated struct {
	EventMetaData
	GUID        string `json:"guid"`
	Name        string `json:"name"`
	ProjectGUID string `json:"ProjectGuid"`
}

func (e *GameCreated) EventType() EventType {
	return GameCreatedType
}
