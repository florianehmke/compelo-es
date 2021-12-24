package compelo_test

import (
	"compelo/command"
	"compelo/event"
	"compelo/query"
	"log"
	"testing"
	"time"
)

func Test(t *testing.T) {
	log.Println("Starting Test")

	bus := event.NewEventBus()
	query.New(bus)
	store := event.CreateEventStore(bus)
	events := store.LoadEvents()

	command := command.New(store, events)

	command.CreateNew("Test")

	time.Sleep(3)
}
