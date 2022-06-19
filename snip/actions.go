package snip

import "fmt"

type SnakesAndLadders struct {
	DiceRoller
	GameGenerator
	game     Game
	finished bool
	turn     int
}

type Actions interface {
	Start()
	move()
}

func (snl *SnakesAndLadders) move() {
	dice := snl.roll()
	loc := snl.game.locations[snl.turn]
	for dice+loc > snl.game.board.size {
		dice = snl.roll()
	}
	if dice+loc == snl.game.board.size {
		snl.finished = true
	}
	nx := dice + loc
	fmt.Printf("moving from %d to %d\n", loc, nx)
	if v, ok := snl.game.board.moves[nx]; ok {
		fmt.Printf("moving from %d to %d\n", nx, v)
		nx = v
	}
	snl.game.locations[snl.turn] = nx
	snl.turn = (snl.turn + 1) % len(snl.game.players)
}
