package event

import (
	"log"
	"sync"
)

type EventBus struct {
	mu     sync.RWMutex
	subs   []chan Event
	closed bool
}

func NewEventBus() *EventBus {
	bus := &EventBus{}
	bus.subs = []chan Event{}
	return bus
}

func (bus *EventBus) Publish(event Event) {
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

func (bus *EventBus) Subscribe() <-chan Event {
	bus.mu.Lock()
	defer bus.mu.Unlock()

	ch := make(chan Event, 100)
	bus.subs = append(bus.subs, ch)
	return ch
}

func (bus *EventBus) Close() {
	bus.mu.Lock()
	defer bus.mu.Unlock()

	if !bus.closed {
		bus.closed = true
		for _, ch := range bus.subs {
			close(ch)
		}
	}
}
