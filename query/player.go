package query

type Player struct {
	GUID        string `json:"guid"`
	ProjectGUID string `json:"projectGuid"`

	Name string `json:"name"`

	ratings map[string]Rating
}

func (p *Player) updateRating(gameGUID string, rating Rating) {
	p.ratings[gameGUID] = rating
}
