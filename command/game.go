package command

type game struct {
	guid        string
	projectGUID string

	name string

	matches map[string]match
}
