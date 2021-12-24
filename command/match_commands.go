package command

import (
	"compelo/event"

	"github.com/google/uuid"
)

func (c *Compelo) CreateNewMatch(gameGUID string, projectGUID string) Response {
	guid := uuid.New().String()
	c.raise(&event.MatchCreated{
		GUID:        guid,
		GameGUID:    gameGUID,
		ProjectGUID: projectGUID,
	})
	return Response{GUID: guid}
}
