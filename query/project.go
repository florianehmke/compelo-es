package query

type Project struct {
	GUID string `json:"guid"`

	Name string `json:"name"`

	players map[string]*Player
	games   map[string]*Game
}
