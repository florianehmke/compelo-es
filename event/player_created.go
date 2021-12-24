package event

import (
	"encoding/json"
	"log"
)

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

func (e *PlayerCreated) UnmarshalFn() func([]byte) Event {
	return func(data []byte) Event {
		var e PlayerCreated
		if err := json.Unmarshal(data, &e); err != nil {
			log.Fatal(err) // TODO: Handle me.
		}
		return &e
	}
}
