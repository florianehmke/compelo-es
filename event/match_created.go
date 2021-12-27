package event

import (
	"time"
)

const MatchCreatedType EventType = "MatchCreated"

// MatchCreated event.
type MatchCreated struct {
	EventMetaData
	GUID        string    `json:"guid"`
	GameGUID    string    `json:"gameGuid"`
	ProjectGUID string    `json:"projectGuid"`
	Date        time.Time `json:"date" ts_type:"string"`
	Teams       []struct {
		PlayerGUIDs []string
		Score       int
	} `json:"teams"`
}

func (e *MatchCreated) EventType() EventType {
	return MatchCreatedType
}
