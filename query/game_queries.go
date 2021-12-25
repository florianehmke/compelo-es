package query

func (c *Compelo) GetAllGames(projectGUID string) []Game {
	c.RLock()
	defer c.RUnlock()

	list := make([]Game, 0, len(c.projects[projectGUID].games))
	for _, value := range c.projects[projectGUID].games {
		list = append(list, value)
	}

	return list
}

func (c *Compelo) GetGameBy(projectGUID string, gameGUID string) Game {
	c.RLock()
	defer c.RUnlock()

	// TODO: Handle not found
	return c.projects[projectGUID].games[gameGUID]
}
