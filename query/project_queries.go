package query

import (
	"errors"
	"fmt"
)

var ErrProjectNotFound = errors.New("project not found")

func (c *Compelo) GetProjects() []*Project {
	c.RLock()
	defer c.RUnlock()

	list := make([]*Project, 0, len(c.projects))
	for _, value := range c.projects {
		list = append(list, value)
	}

	return list
}

func (c *Compelo) GetProjectBy(projectGUID string) (*Project, error) {
	c.RLock()
	defer c.RUnlock()

	return c.getProjectBy(projectGUID)
}

func (c *Compelo) getProjectBy(projectGUID string) (*Project, error) {
	if project, ok := c.projects[projectGUID]; ok {
		return project, nil
	}
	return nil, fmt.Errorf("get project by guid (%s) failed: %w", projectGUID, ErrProjectNotFound)
}
