package compelo_test

import (
	"compelo/command"
	"compelo/event"
	"compelo/query"
	"log"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
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

	time.Sleep(time.Second * 1)

	// Check stuff...
	assert.NotEmpty(t, projectGUID)
	assert.NotEmpty(t, player1GUID)
	assert.NotEmpty(t, player2GUID)
	assert.NotEmpty(t, gameGUID)
	assert.NotEmpty(t, matchGUID)

	assert.Len(t, q.GetProjects(), 2)
	assert.Len(t, q.GetPlayersBy(projectGUID), 2)
	assert.Len(t, q.GetGamesBy(projectGUID), 1)
	assert.Len(t, q.GetMatchesBy(projectGUID, gameGUID), 1)

	project := q.GetProjectBy(projectGUID)
	game := q.GetGameBy(projectGUID, gameGUID)
	player1 := q.GetPlayerBy(projectGUID, player1GUID)
	player2 := q.GetPlayerBy(projectGUID, player2GUID)
	match := q.GetMatchBy(projectGUID, gameGUID, matchGUID)

	assert.NotEmpty(t, project)
	assert.Equal(t, project.Name, "Project 1")

	assert.NotEmpty(t, game)
	assert.Equal(t, game.ProjectGUID, projectGUID)
	assert.Equal(t, game.Name, "Game 1")

	assert.NotEmpty(t, player1)
	assert.Equal(t, player1.ProjectGUID, projectGUID)
	assert.Equal(t, player1.Name, "Player 1")

	assert.NotEmpty(t, player2)
	assert.Equal(t, player2.ProjectGUID, projectGUID)
	assert.Equal(t, player2.Name, "Player 2")

	assert.NotEmpty(t, match)
	assert.Equal(t, match.GameGUID, gameGUID)
	assert.Equal(t, match.ProjectGUID, projectGUID)
	assert.Len(t, match.Teams, 2)
	assert.Len(t, match.Teams[0].Players, 1)
	assert.Len(t, match.Teams[1].Players, 1)

	assert.Equal(t, match.Teams[0].Score, 1)
	assert.Equal(t, match.Teams[0].Result, query.Loss)
	assert.Equal(t, match.Teams[0].RatingDelta, -16)

	assert.Equal(t, match.Teams[1].Score, 2)
	assert.Equal(t, match.Teams[1].Result, query.Win)
	assert.Equal(t, match.Teams[1].RatingDelta, 16)
}
