package command

// Project is a single project in compelo.
type project struct {
	guid string

	name         string
	passwordHash []byte

	players map[string]player
	games   map[string]game
}
