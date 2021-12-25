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

	// Simulate some prior events to ensure re-hydration works.
	store.StoreEvent(&event.ProjectCreated{GUID: "guid", Name: "First Project"})

	// Setup query.
	query := query.New(bus)

	// Load all events from db (rehydrates queries).
	events := store.LoadEvents()

	// Setup command (from existing events).
	command := command.New(store, events)

	// Simulate interaction with command.
	testBasicWorkflow(t, command, query)
}

func testBasicWorkflow(t *testing.T, c *command.Compelo, q *query.Compelo) {
	// 1. Create a project.
	projectGUID := c.CreateNewProject(command.CreateNewProjectCommand{
		Name: "Project 1",
	}).GUID

	// 2. Create two players.
	player1GUID := c.CreateNewPlayer(command.CreateNewPlayerCommand{
		Name:        "Player 1",
		ProjectGUID: projectGUID,
	}).GUID
	player2GUID := c.CreateNewPlayer(command.CreateNewPlayerCommand{
		Name:        "Player 2",
		ProjectGUID: projectGUID,
	}).GUID

	// 3. Create a game.
	gameGUID := c.CreateNewGame(command.CreateNewGameCommand{
		Name:        "Game 1",
		ProjectGUID: projectGUID,
	}).GUID

	// 4. Create a match.
	matchGUID := c.CreateNewMatch(command.CreateNewMatchCommand{
		GameGUID:    gameGUID,
		ProjectGUID: projectGUID,
		Teams: []struct {
			PlayerGUIDs []string
			Score       int
		}{
			{Score: 1, PlayerGUIDs: []string{player1GUID}},
			{Score: 2, PlayerGUIDs: []string{player2GUID}},
		},
	}).GUID

	if projectGUID == "" || player1GUID == "" || player2GUID == "" || gameGUID == "" || matchGUID == "" {
		t.Error("Fatal error...")
	}

	time.Sleep(time.Second * 1)

	projects := q.GetAllProjects()
	if len(projects) != 2 {
		t.Error("Projects in query should be 2")
	}
	log.Println(projects)

	log.Println("Finished!")
}
