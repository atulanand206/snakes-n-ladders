package snip

type Player struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Snake struct {
	Head int `json:"head"`
	Tail int `json:"tail"`
}

type Ladder struct {
	Bottom int `json:"bottom"`
	Top    int `json:"top"`
}

type Board struct {
	Size  int         `json:"size"`
	Moves map[int]int `json:"moves"`
}

type Game struct {
	Board     Board          `json:"board"`
	Locations map[int]int    `json:"locations"`
	Players   map[int]Player `json:"players"`
}
