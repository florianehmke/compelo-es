package query

import "time"

type Match struct {
	GUID        string `json:"guid"`
	GameGUID    string `json:"gameGuid"`
	ProjectGUID string `json:"projectGuid"`

	Date  time.Time `json:"date" ts_type:"string"`
	Teams []Team
}

type Team struct {
	Players []Player `json:"players"`
	Score   int
}
