package command

import (
	"compelo/event"

	"github.com/google/uuid"
)

type CreateNewPlayerCommand struct {
	ProjectGUID string `json:"projectGuid"`
	Name        string `json:"name"`
}

func (c *Compelo) CreateNewPlayer(cmd CreateNewPlayerCommand) Response {
	c.Lock()
	defer c.Unlock()

	// TODO: Check if name is already taken.

	guid := uuid.New().String()
	c.raise(&event.PlayerCreated{
		GUID:        guid,
		ProjectGUID: cmd.ProjectGUID,
		Name:        cmd.Name,
	})
	return Response{GUID: guid}
}
