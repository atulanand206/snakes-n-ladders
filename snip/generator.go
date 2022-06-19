package snip

import "strconv"

type GameGenerator struct {
}

type Generator interface {
	generateGame(size int, players int) (game Game)
}

type PlayerGenerator interface {
	generatePlayers(count int) (players []Player)
	defaultLocations(players []Player) (locations map[int]int)
}

type BoardGenerator interface {
	generateRandom(size int) (board Board)
	generate(size int, snakes []Snake, ladders []Ladder) (board Board)
}

type ObstacleGenerator interface {
	processObstacles(snakes []Snake, ladders []Ladder) (moves map[int]int)
	generateSnakes(count int) (snakes []Snake)
	generateLadders(count int) (ladders []Ladder)
}

func (snl *GameGenerator) generateGame(size int, playersCount int) (game Game) {
	players := snl.generatePlayers((playersCount))
	return Game{
		Board:     snl.generateRandom(size),
		Locations: snl.defaultLocations(players),
		Players:   players,
	}
}

func (snl *GameGenerator) generatePlayers(count int) (players map[int]Player) {
	players = make(map[int]Player)
	i := 0
	for i < count {
		players[i] = Player{i, ("Player" + strconv.Itoa(i))}
		i++
	}
	return
}

func (snl *GameGenerator) defaultLocations(players map[int]Player) (locations map[int]int) {
	locations = make(map[int]int)
	for _, player := range players {
		locations[player.Id] = 1
	}
	return locations
}

func (snl *GameGenerator) generateRandom(size int) (board Board) {
	snakes := snl.generateSnakes(size)
	ladders := snl.generateLadders(size)
	return snl.generate(size, snakes, ladders)
}

func (snl *GameGenerator) generate(size int, snakes []Snake, ladders []Ladder) (board Board) {
	return Board{
		Size:  size * size,
		Moves: snl.processObstacles(snakes, ladders),
	}
}

func (snl *GameGenerator) processObstacles(snakes []Snake, ladders []Ladder) (moves map[int]int) {
	moves = make(map[int]int)
	for _, snake := range snakes {
		moves[snake.Head] = snake.Tail
	}
	for _, ladder := range ladders {
		moves[ladder.Bottom] = ladder.Top
	}
	return
}

func (snl *GameGenerator) generateSnakes(count int) (snakes []Snake) {
	snakes = make([]Snake, count)
	i := 1
	for i < count {
		snakes[i] = Snake{Tail: i * 2, Head: i * 4}
		i++
	}
	return
}

func (snl *GameGenerator) generateLadders(count int) (ladders []Ladder) {
	ladders = make([]Ladder, count)
	i := 1
	for i < count {
		ladders[i] = Ladder{Bottom: i * 3, Top: i * 5}
		i++
	}
	return
}
