package query

type Rating struct {
	PlayerGUID string `json:"playerGuid"`
	GameGUID   string `json:"gameGuid"`
	Current    int    `json:"rating"`
}
