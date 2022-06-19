package snip

import "fmt"

func (snl SnakesAndLadders) Start() {
	snl.game = snl.generateGame(5, 2)
	snl.turn = 0
	snl.finished = false
	fmt.Println(snl.game)
	for {
		if snl.finished {
			fmt.Println("Game is finished.")
			return
		}
		// fmt.Println(snl.turn)
		p := snl.turn
		fmt.Printf("Player %d's move\n", p)
		snl.move()
		fmt.Printf("Player %d's is now at %d\n", p, snl.game.locations[p])
	}
}
