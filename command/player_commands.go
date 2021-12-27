package command

import (
	"compelo/event"
	"fmt"

	"github.com/google/uuid"
)

type CreateNewPlayerCommand struct {
	ProjectGUID string `json:"projectGuid"`
	Name        string `json:"name"`
}

func (c *Compelo) CreateNewPlayer(cmd CreateNewPlayerCommand) (Response, error) {
	c.Lock()
	defer c.Unlock()

	if err := c.checkUniqueConstraint(cmd.ProjectGUID + ":" + cmd.Name); err != nil {
		return Response{}, fmt.Errorf("player name is taken: %w", err)
	}

	guid := uuid.New().String()
	c.raise(&event.PlayerCreated{
		GUID:        guid,
		ProjectGUID: cmd.ProjectGUID,
		Name:        cmd.Name,
	})
	return Response{GUID: guid}, nil
}
