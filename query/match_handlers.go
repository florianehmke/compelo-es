package query

import (
	"compelo/event"
)

func (c *Compelo) handleMatchCreated(e *event.MatchCreated) {
	ratings := make(map[string]int)
	teams := []Team{}

	for _, t := range e.Teams {
		var players []Player
		for _, guid := range t.PlayerGUIDs {
			players = append(players, c.projects[e.ProjectGUID].players[guid])
			ratings[guid] = c.getRatingBy(e.ProjectGUID, guid, e.GameGUID).Current
		}

		teams = append(teams, Team{
			Score:   t.Score,
			Players: players,
		})
	}

	match := Match{
		GUID:        e.GUID,
		GameGUID:    e.GameGUID,
		ProjectGUID: e.ProjectGUID,
		Date:        e.Date,
		Teams:       teams,
	}

	match.determineResult()
	match.calculateTeamElo(ratings)

	for _, team := range teams {
		for _, player := range team.Players {
			player.ratings[e.GameGUID] = Rating{
				PlayerGUID: player.GUID,
				GameGUID:   e.GameGUID,
				Current:    ratings[player.GUID] + team.RatingDelta,
			}
		}
	}

	c.projects[e.ProjectGUID].games[e.GameGUID].matches[e.GUID] = match
}
