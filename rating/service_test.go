package rating_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"compelo/rating"
)

func TestEloForTwoPlayers(t *testing.T) {
	em := rating.NewRatedMatch()

	em.AddPlayer(1, 1, 1500)
	em.AddPlayer(2, 2, 1500)
	em.Calculate()

	assert.Equal(t, 1516, em.GetNewRating(1))
	assert.Equal(t, 16, em.GetRatingDelta(1))

	assert.Equal(t, 1484, em.GetNewRating(2))
	assert.Equal(t, -16, em.GetRatingDelta(2))
}

func TestEloForThreePlayers(t *testing.T) {
	em := rating.NewRatedMatch()

	em.AddPlayer(1, 1, 1500)
	em.AddPlayer(2, 2, 1500)
	em.AddPlayer(3, 3, 1500)
	em.Calculate()

	assert.Equal(t, 1516, em.GetNewRating(1))
	assert.Equal(t, 16, em.GetRatingDelta(1))

	assert.Equal(t, 1500, em.GetNewRating(2))
	assert.Equal(t, 0, em.GetRatingDelta(2))

	assert.Equal(t, 1484, em.GetNewRating(3))
	assert.Equal(t, -16, em.GetRatingDelta(3))
}
