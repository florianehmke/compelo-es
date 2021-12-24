package query

import (
	"compelo/event"
	"log"
)

type Compelo struct {
	projects map[string]Project

	bus *event.Bus
}

func New(bus *event.Bus) *Compelo {
	c := Compelo{
		projects: make(map[string]Project),
		bus:      bus,
	}

	channel := bus.Subscribe()
	go func() {
		for event := range channel {
			c.on(event)
		}
	}()

	return &c
}

type Project struct {
	GUID string `json:"guid"`
	Name string `json:"name"`
}

func (c *Compelo) on(e event.Event) {
	log.Println("Query handling event ", e.GetID(), e.EventType())
	switch e := e.(type) {
	case *event.ProjectCreated:
		c.projects[e.GUID] = Project{
			GUID: e.GUID,
			Name: e.Name,
		}
	case *event.ProjectRenamed:
		if project, ok := c.projects[e.GUID]; ok {
			project.Name = e.NewName
			c.projects[e.GUID] = project
		}
	case *event.ProjectDeleted:
		delete(c.projects, e.GUID)
	}
}

func (p *Compelo) Projects() map[string]Project {
	return p.projects
}
