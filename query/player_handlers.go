package query

import "compelo/event"

func (c *Compelo) handlePlayerCreated(e *event.PlayerCreated) {
	c.projects[e.ProjectGUID].players[e.GUID] = &Player{
		GUID:        e.GUID,
		ProjectGUID: e.ProjectGUID,
		Name:        e.Name,
		ratings:     make(map[string]*Rating),
	}
}
