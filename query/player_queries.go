package query

func (c *Compelo) GetPlayersBy(projectGUID string) []Player {
	c.RLock()
	defer c.RUnlock()

	list := make([]Player, 0, len(c.projects[projectGUID].players))
	for _, value := range c.projects[projectGUID].players {
		list = append(list, *value)
	}

	return list
}

func (c *Compelo) GetPlayerBy(projectGUID string, playerGUID string) Player {
	c.RLock()
	defer c.RUnlock()

	// TODO: Handle not found
	return *c.projects[projectGUID].players[playerGUID]
}
