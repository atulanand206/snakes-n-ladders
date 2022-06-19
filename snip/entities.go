package snip

type Player struct {
	id   int
	name string
}

type Snake struct {
	head int
	tail int
}

type Ladder struct {
	bottom int
	top    int
}

type Board struct {
	size  int
	moves map[int]int
}

type Game struct {
	board     Board
	locations map[int]int
	players   map[int]Player
}
