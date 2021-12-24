package event

import (
	"encoding/json"
	"log"
)

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

func (e *GameCreated) UnmarshalFn() func([]byte) Event {
	return func(data []byte) Event {
		var e GameCreated
		if err := json.Unmarshal(data, &e); err != nil {
			log.Fatal(err) // TODO: Handle me.
		}
		return &e
	}
}
