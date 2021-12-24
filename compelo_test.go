package compelo_test

import (
	"compelo/command"
	"compelo/event"
	"compelo/query"
	"log"
	"os"
	"testing"
	"time"
)

func Test(t *testing.T) {
	log.Println("Starting Test")
	defer os.Remove("test.db")

	// Create event store.
	bus := event.NewBus()
	store := event.NewStore(bus, "test.db")

	// Simulate some prior events.
	store.StoreEvent(&event.ProjectCreated{GUID: "guid", Name: "First Project"})
	store.StoreEvent(&event.ProjectRenamed{GUID: "guid", NewName: "First Project (New Name)"})

	// Setup query.
	query := query.New(bus)

	// Load all events from db (rehydrates queries).
	events := store.LoadEvents()

	// Setup command (from existing events).
	command := command.New(store, events)

	// Simulate interaction with command.
	command.CreateNew("Second Project")

	// Give query time to catch up.
	time.Sleep(time.Second * 1)

	if len(command.Projects()) != 2 {
		t.Error("Projects in command should be 2")
	}

	if len(query.Projects()) != 2 {
		t.Error("Projects in query should be 2")
	}
}
