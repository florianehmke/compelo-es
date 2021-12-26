package query

import (
	"compelo/rating"
	"time"
)

type Result string

const (
	Win  Result = "Win"
	Draw Result = "Draw"
	Loss Result = "Loss"
)

type Match struct {
	GUID        string `json:"guid"`
	GameGUID    string `json:"gameGuid"`
	ProjectGUID string `json:"projectGuid"`

	Date  time.Time `json:"date" ts_type:"string"`
	Teams []Team    `json:"teams"`
}

type Team struct {
	Players     []Player `json:"players"`
	Score       int      `json:"score"`
	Result      Result   `json:"result"`
	RatingDelta int      `json:"ratingDelta"`
}

func (m *Match) determineResult() {
	highScore := 0
	highScoreCount := 0
	for _, t := range m.Teams {
		if t.Score > highScore {
			highScore = t.Score
			highScoreCount = 1
		} else if t.Score == highScore {
			highScoreCount += 1
		}
	}
	if highScoreCount < len(m.Teams) {
		for i := range m.Teams {
			if m.Teams[i].Score == highScore {
				m.Teams[i].Result = Win
			} else {
				m.Teams[i].Result = Loss
			}
		}
	} else {
		for i := range m.Teams {
			m.Teams[i].Result = Draw
		}
	}
}

func (m *Match) calculateTeamElo(ratings map[string]int) {
	rm := rating.NewRatedMatch()
	for i, t := range m.Teams {
		sum := 0
		for _, p := range t.Players {
			sum += ratings[p.GUID]
		}
		avg := sum / len(t.Players)

		// The rating service expects a "rank" to sort players.
		// Here we just use the negative score instead, should
		// result in the same thing for most games..
		rm.AddPlayer(i, -t.Score, avg)
	}
	rm.Calculate()

	for i := range m.Teams {
		m.Teams[i].RatingDelta = rm.GetRatingDelta(i)
	}
}
