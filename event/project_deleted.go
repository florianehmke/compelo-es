package event

import (
	"encoding/json"
	"log"
)

const ProjectDeletedType EventType = "ProjectDeleted"

// ProjectDeleted event.
type ProjectDeleted struct {
	EventMetaData
	GUID string `json:"guid"`
}

func (e *ProjectDeleted) EventType() EventType {
	return ProjectDeletedType
}

func (e *ProjectDeleted) UnmarshalFn() func([]byte) Event {
	return func(data []byte) Event {
		var e ProjectDeleted
		if err := json.Unmarshal(data, &e); err != nil {
			log.Fatal(err)
		}
		return &e
	}
}
