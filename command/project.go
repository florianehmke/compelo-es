package command

// Project is a single project in compelo.
type Project struct {
	guid         string
	name         string
	passwordHash []byte

	players map[string]Player
	games   map[string]game
}

func (c *Compelo) Projects() map[string]Project {
	return c.projects
}
