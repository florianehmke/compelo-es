package compelo_test

import (
	"compelo/command"
	"compelo/event"
	"compelo/query"
	"log"
	"os"
	"testing"
)

func Test(t *testing.T) {
	log.Println("Starting Test")
	defer os.Remove("test.db")

	// Create event store.
	bus := event.NewBus()
	store := event.NewStore(bus, "test.db")

	// Simulate some prior events to ensure re-hydration works.
	store.StoreEvent(&event.ProjectCreated{GUID: "guid", Name: "First Project"})

	// Setup query.
	query.New(bus)

	// Load all events from db (rehydrates queries).
	events := store.LoadEvents()

	// Setup command (from existing events).
	command := command.New(store, events)

	// Simulate interaction with command.
	testBasicWorkflow(t, command)
}

func testBasicWorkflow(t *testing.T, c *command.Compelo) {
	projectGUID := c.CreateNewProject("Second Project").GUID

	player1GUID := c.CreateNewPlayer(projectGUID, "Player 1").GUID
	player2GUID := c.CreateNewPlayer(projectGUID, "Player 2").GUID

	gameGUID := c.CreateNewGame(projectGUID, "Game 1").GUID

	matchGUID := c.CreateNewMatch(gameGUID, projectGUID).GUID

	if projectGUID == "" || player1GUID == "" || player2GUID == "" || gameGUID == "" || matchGUID == "" {
		t.Error("Projects in command should be 2")
	}
	log.Println("Finished!")
}
