package command

import (
	"compelo/event"

	"github.com/google/uuid"
)

func (c *Compelo) CreateNewProject(name string) Response {
	// TODO: Check if name is already taken.

	guid := uuid.New().String()
	c.raise(&event.ProjectCreated{GUID: guid, Name: name})
	return Response{GUID: guid}
}
