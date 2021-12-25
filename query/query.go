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

func (c *Compelo) on(e event.Event) {
	c.Lock()
	defer c.Unlock()

	log.Println("Query handling event ", e.GetID(), e.EventType())

	switch e := e.(type) {
	case *event.ProjectCreated:
		c.handleProjectCreated(e)
	case *event.GameCreated:
		c.handleGameCreated(e)
	case *event.PlayerCreated:
		c.handlePlayerCreated(e)
	}
}
