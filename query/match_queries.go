package query

func (c *Compelo) GetMatchesBy(projectGUID string, gameGUID string) []Match {
	c.RLock()
	defer c.RUnlock()

	list := make([]Match, 0, len(c.projects[projectGUID].games[gameGUID].matches))
	for _, value := range c.projects[projectGUID].games[gameGUID].matches {
		list = append(list, *value)
	}

	return list
}

func (c *Compelo) GetMatchBy(projectGUID string, gameGUID string, matchGUID string) Match {
	c.RLock()
	defer c.RUnlock()

	// TODO: Handle not found
	return *c.projects[projectGUID].games[gameGUID].matches[matchGUID]
}
