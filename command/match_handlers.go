package command

import "compelo/event"

func (c *Compelo) handleMatchCreated(e *event.MatchCreated) {

	// Map teams to interal type.
	var teams []team
	for _, t := range e.Teams {
		teams = append(teams, team{
			playerGUIDs: t.PlayerGUIDs,
			score:       t.Score,
		})
	}

	c.projects[e.ProjectGUID].games[e.GameGUID].matches[e.GUID] = match{
		guid:        e.GUID,
		gameGUID:    e.GameGUID,
		projectGUID: e.ProjectGUID,
		teams:       teams,
	}
}
