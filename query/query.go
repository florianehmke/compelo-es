package query

import (
	"compelo/event"
	"log"
	"sync"
)

type Compelo struct {
	projects map[string]Project

	sync.RWMutex
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
	c.Lock()
	defer c.Unlock()

	log.Println("Query handling event ", e.GetID(), e.EventType())

	switch e := e.(type) {
	case *event.ProjectCreated:
		c.projects[e.GUID] = Project{
			GUID: e.GUID,
			Name: e.Name,
		}
	}
}

func (p *Compelo) Projects() map[string]Project {
	return p.projects
}
