package command

type game struct {
	guid        string
	name        string
	projectGUID string

	matches map[string]match
}
