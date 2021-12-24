package command

import "compelo/event"

func (c *Compelo) handleProjectCreated(e *event.ProjectCreated) {
	c.projects[e.GUID] = Project{
		guid:    e.GUID,
		name:    e.Name,
		games:   make(map[string]game),
		players: make(map[string]Player),
	}
}
