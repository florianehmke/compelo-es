package query

import (
	"errors"
	"fmt"
)

var ErrGameNotFound = errors.New("game not found")

func (c *Compelo) GetGamesBy(projectGUID string) ([]*Game, error) {
	c.RLock()
	defer c.RUnlock()

	project, err := c.getProjectBy(projectGUID)
	if err != nil {
		return nil, fmt.Errorf("get games failed: %w", err)
	}

	list := make([]*Game, 0, len(project.games))
	for _, value := range c.projects[projectGUID].games {
		list = append(list, value)
	}

	return list, nil
}

func (c *Compelo) GetGameBy(projectGUID string, gameGUID string) (*Game, error) {
	c.RLock()
	defer c.RUnlock()

	return c.getGameBy(projectGUID, gameGUID)
}

func (c *Compelo) getGameBy(projectGUID string, gameGUID string) (*Game, error) {
	project, err := c.getProjectBy(projectGUID)
	if err != nil {
		return nil, fmt.Errorf("get game failed: %w", err)
	}

	if game, ok := project.games[gameGUID]; ok {
		return game, nil
	} else {
		return nil, fmt.Errorf("get game by guid (%s) failed: %w", gameGUID, ErrGameNotFound)
	}
}
