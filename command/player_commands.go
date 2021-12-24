package command

import (
	"compelo/event"

	"github.com/google/uuid"
)

func (c *Compelo) CreateNewPlayer(projectGUID string, name string) Response {
	// TODO: Check if name is already taken.

	guid := uuid.New().String()
	c.raise(&event.PlayerCreated{
		GUID:        guid,
		ProjectGUID: projectGUID,
		Name:        name,
	})
	return Response{GUID: guid}
}
