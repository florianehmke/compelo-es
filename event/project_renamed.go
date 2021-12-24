package event

import (
	"encoding/json"
	"log"
)

const ProjectRenamedType EventType = "ProjectRenamed"

// ProjectRenamed event.
type ProjectRenamed struct {
	EventMetaData
	GUID    string `json:"guid"`
	NewName string `json:"new_name"`
}

func (e *ProjectRenamed) EventType() EventType {
	return ProjectRenamedType
}

func (e *ProjectRenamed) UnmarshalFn() func([]byte) Event {
	return func(data []byte) Event {
		var e ProjectRenamed
		if err := json.Unmarshal(data, &e); err != nil {
			log.Fatal(err) // TODO: Handle me.
		}
		return &e
	}
}
