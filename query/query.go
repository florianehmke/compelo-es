package query

import (
	"compelo/event"
	"log"
	"sync"
)

type Handler interface {
	on(e event.Event)
}

type Compelo struct {
	projects map[string]Project

	handlers []Handler

	sync.RWMutex
	bus *event.Bus
}

func New(bus *event.Bus) *Compelo {
	c := Compelo{
		projects: make(map[string]Project),
		bus:      bus,
	}

	c.handlers = []Handler{&RatingHandler{&c}}

	channel := bus.Subscribe()
	go func() {
		for event := range channel {
			c.on(event)

			// Send event to other handlers aswell.
			for _, h := range c.handlers {
				h.on(event)
			}
		}
	}()

	return &c
}

func (c *Compelo) on(e event.Event) {
	c.Lock()
	defer c.Unlock()

	log.Println("[Root Handler] Handling event ", e.GetID(), e.EventType())

	switch e := e.(type) {
	case *event.ProjectCreated:
		c.handleProjectCreated(e)
	case *event.GameCreated:
		c.handleGameCreated(e)
	case *event.PlayerCreated:
		c.handlePlayerCreated(e)
	case *event.MatchCreated:
		c.handleMatchCreated(e)
	}
}
