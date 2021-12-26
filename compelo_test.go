package compelo_test

import (
	"compelo/command"
	"compelo/event"
	"compelo/query"
	"errors"
	"log"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type basicProject struct {
	projectName string
	projectGuid string
	gameName    string
	gameGuid    string
	player1Name string
	player1Guid string
	player2Name string
	player2Guid string

	matchGuid    string
	player1Score int
	player2Score int
}

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
	var testProject = basicProject{
		projectName: "Project 1",
		gameName:    "Game 1",
		player1Name: "Player 1",
		player2Name: "Player 2",

		player1Score: 1,
		player2Score: 2,
	}

	// 1. Create a project.
	testProject.projectGuid = c.CreateNewProject(command.CreateNewProjectCommand{
		Name: testProject.projectName,
	}).GUID

	// 2. Create two players.
	testProject.player1Guid = c.CreateNewPlayer(command.CreateNewPlayerCommand{
		Name:        testProject.player1Name,
		ProjectGUID: testProject.projectGuid,
	}).GUID
	testProject.player2Guid = c.CreateNewPlayer(command.CreateNewPlayerCommand{
		Name:        testProject.player2Name,
		ProjectGUID: testProject.projectGuid,
	}).GUID

	// 3. Create a game.
	testProject.gameGuid = c.CreateNewGame(command.CreateNewGameCommand{
		Name:        testProject.gameName,
		ProjectGUID: testProject.projectGuid,
	}).GUID

	// 4. Create a match.
	testProject.matchGuid = c.CreateNewMatch(command.CreateNewMatchCommand{
		GameGUID:    testProject.gameGuid,
		ProjectGUID: testProject.projectGuid,
		Teams: []struct {
			PlayerGUIDs []string
			Score       int
		}{
			{Score: testProject.player1Score, PlayerGUIDs: []string{testProject.player1Guid}},
			{Score: testProject.player2Score, PlayerGUIDs: []string{testProject.player1Guid}},
		},
	}).GUID

	time.Sleep(time.Second * 1)

	checkCommandResults(t, testProject)
	checkQuery(t, q, testProject)

}

func checkCommandResults(t *testing.T, testProject basicProject) {
	assert.NotEmpty(t, testProject.projectGuid)
	assert.NotEmpty(t, testProject.player1Guid)
	assert.NotEmpty(t, testProject.player2Guid)
	assert.NotEmpty(t, testProject.gameGuid)
	assert.NotEmpty(t, testProject.matchGuid)
}

func checkQuery(t *testing.T, q *query.Compelo, testProject basicProject) {
	checkQueryGetProjects(t, q, testProject)
	checkQueryGetPlayersBy(t, q, testProject)
	checkQueryGetGamesBy(t, q, testProject)
	checkQueryGetMatchesBy(t, q, testProject)

	checkQueryGetProjectBy(t, q, testProject)
	checkQueryGetGameBy(t, q, testProject)
	// player1 := q.GetPlayerBy(projectGUID, player1GUID)
	// player2 := q.GetPlayerBy(projectGUID, player2GUID)
	// match := q.GetMatchBy(projectGUID, gameGUID, matchGUID)
	// ratingPlayer1 := q.GetRatingBy(projectGUID, player1GUID, gameGUID)
	// ratingPlayer2 := q.GetRatingBy(projectGUID, player2GUID, gameGUID)

	// assert.NotEmpty(t, project)
	// assert.Equal(t, project.Name, "Project 1")

	// assert.NotEmpty(t, game)
	// assert.Equal(t, game.ProjectGUID, projectGUID)
	// assert.Equal(t, game.Name, "Game 1")

	// assert.NotEmpty(t, player1)
	// assert.Equal(t, player1.ProjectGUID, projectGUID)
	// assert.Equal(t, player1.Name, "Player 1")

	// assert.NotEmpty(t, player2)
	// assert.Equal(t, player2.ProjectGUID, projectGUID)
	// assert.Equal(t, player2.Name, "Player 2")

	// assert.NotEmpty(t, match)
	// assert.Equal(t, match.GameGUID, gameGUID)
	// assert.Equal(t, match.ProjectGUID, projectGUID)
	// assert.Len(t, match.Teams, 2)
	// assert.Len(t, match.Teams[0].Players, 1)
	// assert.Len(t, match.Teams[1].Players, 1)

	// assert.Equal(t, match.Teams[0].Score, 1)
	// assert.Equal(t, match.Teams[0].Result, query.Loss)
	// assert.Equal(t, match.Teams[0].RatingDelta, -16)

	// assert.Equal(t, match.Teams[1].Score, 2)
	// assert.Equal(t, match.Teams[1].Result, query.Win)
	// assert.Equal(t, match.Teams[1].RatingDelta, 16)

	// assert.Equal(t, 1484, ratingPlayer1.Current)
	// assert.Equal(t, 1516, ratingPlayer2.Current)
}

func checkQueryGetProjects(t *testing.T, q *query.Compelo, testProject basicProject) {
	assert.Len(t, q.GetProjects(), 2)
}

func checkQueryGetPlayersBy(t *testing.T, q *query.Compelo, testProject basicProject) {
	players, err := q.GetPlayersBy(testProject.projectGuid)
	assert.Len(t, players, 2)
	assert.Nil(t, err)

	players, err = q.GetPlayersBy("404")
	assert.Nil(t, players)
	assert.True(t, errors.Is(err, query.ErrProjectNotFound))
}

func checkQueryGetGamesBy(t *testing.T, q *query.Compelo, testProject basicProject) {
	games, err := q.GetGamesBy(testProject.projectGuid)
	assert.Len(t, games, 1)
	assert.Nil(t, err)

	games, err = q.GetGamesBy("404")
	assert.Nil(t, games)
	assert.True(t, errors.Is(err, query.ErrProjectNotFound))
}

func checkQueryGetMatchesBy(t *testing.T, q *query.Compelo, testProject basicProject) {
	matches, err := q.GetMatchesBy(testProject.projectGuid, testProject.gameGuid)
	assert.Len(t, matches, 1)
	assert.Nil(t, err)

	matches, err = q.GetMatchesBy("404", testProject.gameGuid)
	assert.Nil(t, matches)
	assert.True(t, errors.Is(err, query.ErrProjectNotFound))

	matches, err = q.GetMatchesBy(testProject.projectGuid, "404")
	assert.Nil(t, matches)
	assert.True(t, errors.Is(err, query.ErrGameNotFound))
}

func checkQueryGetProjectBy(t *testing.T, q *query.Compelo, testProject basicProject) {
	project, err := q.GetProjectBy(testProject.projectGuid)
	assert.NotNil(t, project)
	assert.Equal(t, testProject.projectName, project.Name)
	assert.Nil(t, err)

	project, err = q.GetProjectBy("404")
	assert.Nil(t, project)
	assert.True(t, errors.Is(err, query.ErrProjectNotFound))
}

func checkQueryGetGameBy(t *testing.T, q *query.Compelo, testProject basicProject) {
	game, err := q.GetGameBy(testProject.projectGuid, testProject.gameGuid)
	assert.NotNil(t, game)
	assert.Equal(t, testProject.gameName, game.Name)
	assert.Nil(t, err)

	game, err = q.GetGameBy("404", testProject.gameGuid)
	assert.Nil(t, game)
	assert.True(t, errors.Is(err, query.ErrProjectNotFound))

	game, err = q.GetGameBy(testProject.projectGuid, "404")
	assert.Nil(t, game)
	assert.True(t, errors.Is(err, query.ErrGameNotFound))
}
