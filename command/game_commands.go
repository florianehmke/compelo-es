package command

import (
	"compelo/event"

	"github.com/google/uuid"
)

type CreateNewGameCommand struct {
	ProjectGUID string `json:"projectGuid"`
	Name        string `json:"name"`
}

func (c *Compelo) CreateNewGame(cmd CreateNewGameCommand) Response {
	// TODO: Check if name is already taken.

	guid := uuid.New().String()
	c.raise(&event.GameCreated{
		GUID:        guid,
		ProjectGUID: cmd.ProjectGUID,
		Name:        cmd.Name,
	})
	return Response{GUID: guid}
}
