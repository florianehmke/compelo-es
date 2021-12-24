package command

import (
	"compelo/event"

	"github.com/google/uuid"
)

func (c *Compelo) CreateNewGame(projectGUID string, name string) Response {
	// TODO: Check if name is already taken.

	guid := uuid.New().String()
	c.raise(&event.GameCreated{
		GUID:        guid,
		ProjectGUID: projectGUID,
		Name:        name,
	})
	return Response{GUID: guid}
}
