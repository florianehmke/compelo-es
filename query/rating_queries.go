package query

func (c *Compelo) GetRatingBy(projectGUID string, playerGUID string, gameGUID string) Rating {
	c.RLock()
	defer c.RUnlock()

	return c.getRatingBy(projectGUID, playerGUID, gameGUID)
}

func (c *Compelo) getRatingBy(projectGUID string, playerGUID string, gameGUID string) Rating {
	// TODO: Handle not found
	return c.projects[projectGUID].players[playerGUID].ratings[gameGUID]
}
