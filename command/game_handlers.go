package command

import "compelo/event"

func (c *Compelo) handleGameCreated(e *event.GameCreated) {
	c.projects[e.ProjectGUID].games[e.GUID] = game{
		guid:        e.GUID,
		projectGUID: e.ProjectGUID,
		name:        e.Name,
		matches:     make(map[string]match),
	}
}
