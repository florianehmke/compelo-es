package event

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
