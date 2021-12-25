package event

import (
	"encoding/json"
	"log"
)

const ProjectCreatedType EventType = "ProjectCreated"

// ProjectCreated event.
type ProjectCreated struct {
	EventMetaData
	GUID         string `json:"guid"`
	Name         string `json:"name"`
	PasswordHash []byte `json:"passwordHash"`
}

func (e *ProjectCreated) EventType() EventType {
	return ProjectCreatedType
}

func (e *ProjectCreated) UnmarshalFn() func([]byte) Event {
	return func(data []byte) Event {
		var e ProjectCreated
		if err := json.Unmarshal(data, &e); err != nil {
			log.Fatal(err) // TODO: Handle me.
		}
		return &e
	}
}
