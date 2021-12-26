package query

import "compelo/rating"

func (c *Compelo) GetRatingBy(projectGUID string, playerGUID string, gameGUID string) Rating {
	c.RLock()
	defer c.RUnlock()

	return *c.getRatingBy(projectGUID, playerGUID, gameGUID)
}

func (c *Compelo) getRatingBy(projectGUID string, playerGUID string, gameGUID string) *Rating {
	// TODO: Handle not found

	r := c.projects[projectGUID].players[playerGUID].ratings[gameGUID]
	if r == nil {
		r = &Rating{
			PlayerGUID: playerGUID,
			GameGUID:   gameGUID,
			Current:    rating.InitialRating,
		}
		c.projects[projectGUID].players[playerGUID].ratings[gameGUID] = r
	}
	return r
}
