package event

import (
	"log"
	"sync"
)

type Bus struct {
	mu     sync.RWMutex
	subs   []chan Event
	closed bool
}

func NewBus() *Bus {
	bus := &Bus{}
	bus.subs = []chan Event{}
	return bus
}

func (bus *Bus) Publish(event Event) {
	bus.mu.RLock()
	defer bus.mu.RUnlock()

	if bus.closed {
		return
	}

	for _, ch := range bus.subs {
		log.Println("Publishing", event.EventType())
		ch <- event
	}
}

func (bus *Bus) Subscribe() <-chan Event {
	bus.mu.Lock()
	defer bus.mu.Unlock()

	// TODO: This is far from optimal.
	ch := make(chan Event, 100)
	bus.subs = append(bus.subs, ch)
	return ch
}

func (bus *Bus) Close() {
	bus.mu.Lock()
	defer bus.mu.Unlock()

	if !bus.closed {
		bus.closed = true
		for _, ch := range bus.subs {
			close(ch)
		}
	}
}
