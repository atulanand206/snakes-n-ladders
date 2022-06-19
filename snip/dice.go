package snip

import (
	"math/rand"
	"time"
)

type Roller interface {
	roll()
}

type DiceRoller struct {
}

func (roller *DiceRoller) roll() int {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Intn(6) + 1
}
