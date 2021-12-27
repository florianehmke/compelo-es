package event

import (
	"encoding/json"
	"errors"
	"fmt"
)

var ErrEventUnmarshalFailed = errors.New("unmarshal of event failed")

type EventType string

// Event is a domain event marker.
type Event interface {
	EventType() EventType
	SetID(uint64)
	GetID() uint64
}

func (et EventType) Unmarshal(data []byte) (Event, error) {
	var target Event
	switch et {
	case PlayerCreatedType:
		target = &PlayerCreated{}
	case GameCreatedType:
		target = &GameCreated{}
	case ProjectCreatedType:
		target = &ProjectCreated{}
	case MatchCreatedType:
		target = &MatchCreated{}
	}

	if err := json.Unmarshal(data, &target); err != nil {
		return nil, fmt.Errorf("unmarshal failed: %w", err)
	}

	return target, nil
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
