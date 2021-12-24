package event

import (
	"encoding/json"
	"log"
)

const MatchCreatedType EventType = "MatchCreated"

// MatchCreated event.
type MatchCreated struct {
	EventMetaData
	GUID        string `json:"guid"`
	GameGUID    string `json:"gameGuid"`
	ProjectGUID string `json:"projectGuid"`
	Teams       []struct {
		PlayerGUIDs []string
		Score       int
	} `json:"teams"`
}

func (e *MatchCreated) EventType() EventType {
	return MatchCreatedType
}

func (e *MatchCreated) UnmarshalFn() func([]byte) Event {
	return func(data []byte) Event {
		var e MatchCreated
		if err := json.Unmarshal(data, &e); err != nil {
			log.Fatal(err) // TODO: Handle me.
		}
		return &e
	}
}
