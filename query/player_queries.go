package query

import (
	"errors"
	"fmt"
)

var ErrPlayerNotFound = errors.New("player not found")

func (c *Compelo) GetPlayersBy(projectGUID string) ([]*Player, error) {
	c.RLock()
	defer c.RUnlock()

	project, err := c.getProjectBy(projectGUID)
	if err != nil {
		return nil, fmt.Errorf("get players failed: %w", err)
	}

	list := make([]*Player, 0, len(project.players))
	for _, value := range project.players {
		list = append(list, value)
	}

	return list, nil
}

func (c *Compelo) GetPlayerBy(projectGUID string, playerGUID string) (*Player, error) {
	c.RLock()
	defer c.RUnlock()

	project, err := c.getProjectBy(projectGUID)
	if err != nil {
		return nil, fmt.Errorf("get player failed: %w", err)
	}

	if player, ok := project.players[playerGUID]; ok {
		return player, nil
	} else {
		return nil, fmt.Errorf("get player by guid (%s) failed: %w", playerGUID, ErrPlayerNotFound)
	}
}
