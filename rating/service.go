package rating

import (
	"math"
)

const InitialRating = 1500

type RatedMatch struct {
	players []ratedPlayer
}

func NewRatedMatch() *RatedMatch {
	return &RatedMatch{}
}

type ratedPlayer struct {
	id int

	place       int
	ratingPre   int
	ratingPost  int
	ratingDelta int
}

func (m *RatedMatch) AddPlayer(id, place, elo int) {
	m.players = append(m.players, ratedPlayer{
		id:        id,
		place:     place,
		ratingPre: elo,
	})
}

func (m *RatedMatch) GetNewRating(id int) int {
	return m.findPlayer(id).ratingPost
}

func (m *RatedMatch) GetRatingDelta(id int) int {
	return m.findPlayer(id).ratingDelta
}

func (m *RatedMatch) findPlayer(id int) ratedPlayer {
	for i := 0; i < len(m.players); i++ {
		if m.players[i].id == id {
			return m.players[i]
		}
	}
	return ratedPlayer{}
}

func (m *RatedMatch) Calculate() {
	n := len(m.players)
	k := 32 / (float64)(n-1)

	for i := 0; i < n; i++ {
		place := m.players[i].place
		rating := m.players[i].ratingPre

		for j := 0; j < n; j++ {
			if i == j {
				continue
			}

			opponentPlace := m.players[j].place
			opponentRating := m.players[j].ratingPre

			s := 0.5
			if place < opponentPlace {
				s = 1.0
			}
			if place > opponentPlace {
				s = 0.0
			}

			ea := 1 / (1.0 + math.Pow(10.0, (float64(opponentRating-rating))/400.0))
			m.players[i].ratingDelta += int(math.Round(k * (s - ea)))
		}

		m.players[i].ratingPost = m.players[i].ratingPre + m.players[i].ratingDelta
	}
}
