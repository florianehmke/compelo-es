package command

type match struct {
	guid        string
	gameGUID    string
	projectGUID string
	teams       []team
}

type team struct {
	playerGUIDs []string
	score       int
}
