package query

import (
	"errors"
	"fmt"
)

var ErrMatchNotFound = errors.New("match not found")

func (c *Compelo) GetMatchesBy(projectGUID string, gameGUID string) ([]*Match, error) {
	c.RLock()
	defer c.RUnlock()

	game, err := c.getGameBy(projectGUID, gameGUID)
	if err != nil {
		return nil, fmt.Errorf("get matches failed: %w", err)
	}

	list := make([]*Match, 0, len(game.matches))
	for _, value := range game.matches {
		list = append(list, value)
	}

	return list, nil
}

func (c *Compelo) GetMatchBy(projectGUID string, gameGUID string, matchGUID string) (*Match, error) {
	c.RLock()
	defer c.RUnlock()

	game, err := c.getGameBy(projectGUID, gameGUID)
	if err != nil {
		return nil, fmt.Errorf("get matches failed: %w", err)
	}

	if match, ok := game.matches[matchGUID]; ok {
		return match, nil
	} else {
		return nil, fmt.Errorf("get match by guid (%s) failed: %w", matchGUID, ErrMatchNotFound)
	}
}
