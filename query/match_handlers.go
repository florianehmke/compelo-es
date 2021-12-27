package query

import (
	"compelo/event"
	"log"
)

func (c *Compelo) handleMatchCreated(e *event.MatchCreated) {
	ratings := make(map[string]*Rating)
	teams := []*Team{}

	for _, t := range e.Teams {
		var players []*Player
		for _, guid := range t.PlayerGUIDs {
			players = append(players, c.projects[e.ProjectGUID].players[guid])

			if rating, err := c.getRatingBy(e.ProjectGUID, guid, e.GameGUID); err == nil {
				ratings[guid] = rating
			} else {
				log.Fatalf("unexpected error in handler: %s", err.Error())
			}
		}

		teams = append(teams, &Team{
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
	match.updatePlayerRatings(ratings)

	c.projects[e.ProjectGUID].games[e.GameGUID].matches[e.GUID] = &match
}
