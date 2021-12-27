package command

import (
	"compelo/event"
	"fmt"

	"github.com/google/uuid"
)

type CreateNewGameCommand struct {
	ProjectGUID string `json:"projectGuid"`
	Name        string `json:"name"`
}

func (c *Compelo) CreateNewGame(cmd CreateNewGameCommand) (Response, error) {
	c.Lock()
	defer c.Unlock()

	if err := c.checkUniqueConstraint(cmd.ProjectGUID + ":" + cmd.Name); err != nil {
		return Response{}, fmt.Errorf("game name is taken: %w", err)
	}

	guid := uuid.New().String()
	c.raise(&event.GameCreated{
		GUID:        guid,
		ProjectGUID: cmd.ProjectGUID,
		Name:        cmd.Name,
	})
	return Response{GUID: guid}, nil
}
