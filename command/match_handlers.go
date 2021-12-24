package command

import "compelo/event"

func (c *Compelo) handleMatchCreated(e *event.MatchCreated) {
	c.projects[e.ProjectGUID].games[e.GameGUID].matches[e.GUID] = match{
		guid:        e.GUID,
		gameGUID:    e.GameGUID,
		projectGUID: e.ProjectGUID,
	}
}
