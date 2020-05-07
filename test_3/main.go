package main

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

var BOARD = [][]string{
	{"#", "#", "#", "#", "#", "#", "#", "#"},
	{"#", ".", ".", ".", ".", ".", ".", "#"},
	{"#", ".", "#", "#", "#", ".", ".", "#"},
	{"#", ".", ".", ".", "#", ".", "#", "#"},
	{"#", "X", "#", ".", ".", ".", ".", "#"},
	{"#", "#", "#", "#", "#", "#", "#", "#"},
}

type Location struct {
	x int
	y int
}

type Player struct {
	location Location
	steps    []StepConstraint
}

func NewPlayer(x, y int, steps []StepConstraint) Player {
	return Player{
		location: Location{
			x: x,
			y: y,
		},
		steps: steps,
	}
}

func (p *Player) toNorth() {
	p.location.x -= 1
}

func (p *Player) toEast() {
	p.location.y += 1
}

func (p *Player) toSouth() {
	p.location.x += 1
}

func (p *Player) toWest() {
	p.location.y -= 1
}

func (p *Player) isAllowedToMove() bool {
	floor := BOARD[p.location.x][p.location.y]
	if floor == "#" {
		return false
	}

	return true
}

type StepConstraint struct {
	name   string
	action func(*Player)
	undo   func(*Player)
}

var ToNorth = StepConstraint{
	name: "toNorth",
	action: func(p *Player) {
		p.toNorth()
	},
	undo: func(p *Player) {
		p.toSouth()
	},
}

var ToEast = StepConstraint{
	name: "toEast",
	action: func(p *Player) {
		p.toEast()
	},
	undo: func(p *Player) {
		p.toWest()
	},
}

var ToSouth = StepConstraint{
	name: "toSouth",
	action: func(p *Player) {
		p.toSouth()
	},
	undo: func(p *Player) {
		p.toNorth()
	},
}

func main() {
	steps := []StepConstraint{ToNorth, ToEast, ToSouth}
	player := NewPlayer(4, 1, steps)
	totalStep := countTotalStep(player)
	fmt.Println("Total possibility dot:", totalStep)
}

func countTotalStep(player Player) int {
	total := 0
	for i, step := range player.steps {
		for {
			x := player.location.x
			y := player.location.y
			BOARD[x][y] = "X"
			step.action(&player)
			if !player.isAllowedToMove() {
				step.undo(&player)
				break
			}

			if i < len(player.steps)-1 {
				tempPlayer := NewPlayer(x, y, player.steps[i+1:])
				step := tempPlayer.steps[0]
				step.action(&tempPlayer)
				if tempPlayer.isAllowedToMove() {
					total += countTotalStep(tempPlayer)
				}
			}

			BOARD[player.location.x][player.location.y] = "X"
			total += 1
			displayBoard(BOARD)
		}
	}

	return total
}

func displayBoard(board [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	for _, row := range board {
		table.Append(row)
	}

	table.Render()
}
