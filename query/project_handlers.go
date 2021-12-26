package query

import "compelo/event"

func (c *Compelo) handleProjectCreated(e *event.ProjectCreated) {
	c.projects[e.GUID] = Project{
		GUID:    e.GUID,
		Name:    e.Name,
		games:   make(map[string]*Game),
		players: make(map[string]*Player),
	}
}
