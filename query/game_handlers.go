package query

import "compelo/event"

func (c *Compelo) handleGameCreated(e *event.GameCreated) {
	c.projects[e.ProjectGUID].games[e.GUID] = Game{
		GUID:        e.GUID,
		ProjectGUID: e.ProjectGUID,
		Name:        e.Name,
	}
}
