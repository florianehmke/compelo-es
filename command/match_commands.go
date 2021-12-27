package command

import (
	"compelo/event"
	"time"

	"github.com/google/uuid"
)

type CreateNewMatchCommand struct {
	GameGUID    string `json:"gameGuid"`
	ProjectGUID string `json:"projectGuid"`
	Teams       []struct {
		PlayerGUIDs []string
		Score       int
	} `json:"teams"`
}

func (c *Compelo) CreateNewMatch(cmd CreateNewMatchCommand) (Response, error) {
	c.Lock()
	defer c.Unlock()

	// TODO: validate event

	guid := uuid.New().String()
	c.raise(&event.MatchCreated{
		GUID:        guid,
		GameGUID:    cmd.GameGUID,
		ProjectGUID: cmd.ProjectGUID,
		Date:        time.Now(),
		Teams:       cmd.Teams,
	})
	return Response{GUID: guid}, nil
}
