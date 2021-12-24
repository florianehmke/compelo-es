package command

import (
	"compelo/event"

	"github.com/google/uuid"
)

type CreateNewProjectCommand struct {
	Name string `json:"name"`
}

func (c *Compelo) CreateNewProject(cmd CreateNewProjectCommand) Response {
	// TODO: Check if name is already taken.

	guid := uuid.New().String()
	c.raise(&event.ProjectCreated{GUID: guid, Name: cmd.Name})
	return Response{GUID: guid}
}
