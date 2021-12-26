package query

import (
	"fmt"
)

func (c *Compelo) GetRatingBy(projectGUID string, playerGUID string, gameGUID string) (*Rating, error) {
	c.RLock()
	defer c.RUnlock()

	return c.getRatingBy(projectGUID, playerGUID, gameGUID)
}

func (c *Compelo) getRatingBy(projectGUID string, playerGUID string, gameGUID string) (*Rating, error) {
	player, err := c.getPlayerBy(projectGUID, playerGUID)
	if err != nil {
		return nil, fmt.Errorf("get rating failed: %w", err)
	}

	if r, ok := player.ratings[gameGUID]; ok {
		return r, nil
	} else {
		player.ratings[gameGUID] = initialRatingFor(playerGUID, gameGUID)
		return player.ratings[gameGUID], nil
	}
}
