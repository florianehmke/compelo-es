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
	players     []*basicPlayer
	matchGuid   string
}

type basicPlayer struct {
	guid  string
	name  string
	score int
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
		players: []*basicPlayer{
			{name: "Player 1", score: 1},
			{name: "Player 2", score: 2},
		},
	}

	// 1. Create a project.
	testProject.projectGuid = c.CreateNewProject(command.CreateNewProjectCommand{
		Name: testProject.projectName,
	}).GUID

	// 2. Create two players.
	for _, p := range testProject.players {
		p.guid = c.CreateNewPlayer(command.CreateNewPlayerCommand{
			Name:        p.name,
			ProjectGUID: testProject.projectGuid,
		}).GUID
	}

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
			{Score: testProject.players[0].score, PlayerGUIDs: []string{testProject.players[0].guid}},
			{Score: testProject.players[1].score, PlayerGUIDs: []string{testProject.players[1].guid}},
		},
	}).GUID

	time.Sleep(time.Second * 1)

	checkCommandResults(t, testProject)
	checkQuery(t, q, testProject)

}

func checkCommandResults(t *testing.T, testProject basicProject) {
	assert.NotEmpty(t, testProject.projectGuid)
	assert.NotEmpty(t, testProject.players[0].guid)
	assert.NotEmpty(t, testProject.players[1].guid)
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
	checkQueryGetPlayerBy(t, q, testProject)
	checkQueryGetMatchBy(t, q, testProject)
	checkQueryGetRatingBy(t, q, testProject)

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

func checkQueryGetPlayerBy(t *testing.T, q *query.Compelo, testProject basicProject) {
	for _, p := range testProject.players {
		player, err := q.GetPlayerBy(testProject.projectGuid, p.guid)
		assert.NotNil(t, player)
		assert.Equal(t, p.name, player.Name)
		assert.Nil(t, err)

		player, err = q.GetPlayerBy("404", testProject.gameGuid)
		assert.Nil(t, player)
		assert.True(t, errors.Is(err, query.ErrProjectNotFound))

		player, err = q.GetPlayerBy(testProject.projectGuid, "404")
		assert.Nil(t, player)
		assert.True(t, errors.Is(err, query.ErrPlayerNotFound))
	}
}

func checkQueryGetMatchBy(t *testing.T, q *query.Compelo, testProject basicProject) {
	match, err := q.GetMatchBy(testProject.projectGuid, testProject.gameGuid, testProject.matchGuid)
	assert.NotNil(t, match)
	assert.Nil(t, err)

	match, err = q.GetMatchBy("404", testProject.gameGuid, testProject.matchGuid)
	assert.Nil(t, match)
	assert.True(t, errors.Is(err, query.ErrProjectNotFound))

	match, err = q.GetMatchBy(testProject.projectGuid, "404", testProject.matchGuid)
	assert.Nil(t, match)
	assert.True(t, errors.Is(err, query.ErrGameNotFound))

	match, err = q.GetMatchBy(testProject.projectGuid, testProject.gameGuid, "404")
	assert.Nil(t, match)
	assert.True(t, errors.Is(err, query.ErrMatchNotFound))
}

func checkQueryGetRatingBy(t *testing.T, q *query.Compelo, testProject basicProject) {
	for i, p := range testProject.players {
		rating, err := q.GetRatingBy(testProject.projectGuid, p.guid, testProject.gameGuid)
		assert.NotNil(t, rating)
		assert.Nil(t, err)

		assert.Equal(t, p.guid, rating.PlayerGUID)

		if i == 0 {
			assert.Equal(t, 1484, rating.Current)
		} else if i == 1 {
			assert.Equal(t, 1516, rating.Current)
		}

		rating, err = q.GetRatingBy("404", p.guid, testProject.gameGuid)
		assert.Nil(t, rating)
		assert.True(t, errors.Is(err, query.ErrProjectNotFound))

		rating, err = q.GetRatingBy(testProject.projectGuid, "404", testProject.gameGuid)
		assert.Nil(t, rating)
		assert.True(t, errors.Is(err, query.ErrPlayerNotFound))
	}
}
