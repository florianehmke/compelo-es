package query

import "time"

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
	Players     []Player    `json:"players"`
	Score       int         `json:"score"`
	Result      Result      `json:"result"`
	RatingDelta interface{} `json:"ratingDelta"`
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
