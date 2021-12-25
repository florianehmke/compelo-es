package query

func (c *Compelo) GetAllPlayers(projectGUID string) []Player {
	c.RLock()
	defer c.RUnlock()

	list := make([]Player, 0, len(c.projects[projectGUID].players))
	for _, value := range c.projects[projectGUID].players {
		list = append(list, value)
	}

	return list
}
