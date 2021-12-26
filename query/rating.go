package query

import "compelo/rating"

type Rating struct {
	PlayerGUID string `json:"playerGuid"`
	GameGUID   string `json:"gameGuid"`

	Current int `json:"rating"`
}

func initialRatingFor(playerGUID string, gameGUID string) *Rating {
	return &Rating{
		PlayerGUID: playerGUID,
		GameGUID:   gameGUID,
		Current:    rating.InitialRating,
	}
}
