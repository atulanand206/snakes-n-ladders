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
	loc := snl.game.Locations[snl.turn]
	for dice+loc > snl.game.Board.Size {
		dice = snl.roll()
	}
	if dice+loc == snl.game.Board.Size {
		snl.finished = true
	}
	nx := dice + loc
	fmt.Printf("moving from %d to %d\n", loc, nx)
	if v, ok := snl.game.Board.Moves[nx]; ok {
		fmt.Printf("moving from %d to %d\n", nx, v)
		nx = v
	}
	snl.game.Locations[snl.turn] = nx
	snl.turn = (snl.turn + 1) % len(snl.game.Players)
}
