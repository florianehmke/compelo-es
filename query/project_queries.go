package query

func (c *Compelo) GetAllProjects() []Project {
	c.RLock()
	defer c.RUnlock()

	list := make([]Project, 0, len(c.projects))
	for _, value := range c.projects {
		list = append(list, value)
	}

	return list
}

func (c *Compelo) GetProjectBy(projectGUID string) Project {
	c.RLock()
	defer c.RUnlock()

	// TODO: Handle not found
	return c.projects[projectGUID]
}
