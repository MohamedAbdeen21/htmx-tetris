package game

import (
	"log"
	"tetris/game/actions"
)

const HEIGHT = 15
const WIDTH = 8

type Row [WIDTH]string

func NewRow() Row {
	r := Row{}
	for i := 0; i < WIDTH; i++ {
		r[i] = "E"
	}
	return r
}

type Game struct {
	State              [HEIGHT]Row
	GameOver           bool
	CurrentTetrominoes *Tetrominoes
	Score              int
}

func NewGame() *Game {
	var state [HEIGHT]Row
	for row := range state {
		state[row] = NewRow()
	}
	return &Game{State: state}
}

func (g *Game) Restart() {
	if g.GameOver == false {
		return
	}

	g.State = NewGame().State
	g.CurrentTetrominoes = nil
	g.GameOver = false
	g.Score = 0
}

func (g *Game) Tick(action actions.Action) {
	if g.GameOver {
		return
	}

	if g.CurrentTetrominoes == nil {
		g.CurrentTetrominoes = NewTetrominoes()
	}

	// remove piece from previous buffer
	for _, pos := range g.CurrentTetrominoes.positions {
		if pos.x < 0 {
			continue // out-of-bounds
		}

		g.State[pos.x][pos.y] = "E"
	}

	collided := false

	switch action {
	case actions.Left:
		g.PieceLeft()
	case actions.Right:
		g.PieceRight()
	case actions.Down:
		collided = g.PieceDown()
	}

	// add piece to next buffer
	for _, pos := range g.CurrentTetrominoes.positions {
		if pos.x < 0 {
			continue // out-of-bounds
		}

		if g.State[pos.x][pos.y] != "E" {
			panic("Should never happen")
		}

		g.State[pos.x][pos.y] = g.CurrentTetrominoes.symbol
	}

	// piece collided, remove to spawn a new piece and check
	// if any rows are completed
	if collided {
		g.checkGameOver()
		g.checkCompletedRows()
		g.CurrentTetrominoes = nil
	}

}

func (g *Game) Display() {
	for _, row := range g.State {
		log.Print(row)
	}

	if g.CurrentTetrominoes == nil {
		return
	}

	for _, pos := range g.CurrentTetrominoes.positions {
		log.Print(pos.x, pos.y)
	}
}

func (g *Game) checkGameOver() {
	for _, pos := range g.CurrentTetrominoes.positions {
		if pos.x < 0 {
			g.GameOver = true
			return
		}
	}
	return
}

func (g *Game) checkCompletedRows() {
	state := g.State
	cleared_rows := 0
	for index, row := range state {
		sum := 0
		for _, block := range row {
			if block != "E" {
				sum += 1
			}
		}

		if sum != WIDTH {
			continue
		}

		cleared_rows++
		g.Score += 100

		for i := index; i >= 1; i-- {
			g.State[i] = g.State[i-1]
		}

		g.State[0] = NewRow()
	}
	bonus := max(cleared_rows-1, 0) * 20
	g.Score += bonus
}

func (g *Game) PieceDown() bool {
	for _, pos := range g.CurrentTetrominoes.positions {
		if pos.x < -1 {
			continue
		}
		if pos.x == HEIGHT-1 || g.State[pos.x+1][pos.y] != "E" {
			return true
		}
	}

	for _, pos := range g.CurrentTetrominoes.positions {
		pos.x++
	}

	return false
}

func (g *Game) PieceLeft() {
	for _, pos := range g.CurrentTetrominoes.positions {
		if pos.y == 0 || pos.x >= 0 && g.State[pos.x][pos.y-1] != "E" {
			return
		}
	}

	for _, pos := range g.CurrentTetrominoes.positions {
		pos.y--
	}
}

func (g *Game) PieceRight() {
	for _, pos := range g.CurrentTetrominoes.positions {
		if pos.y == WIDTH-1 || pos.x >= 0 && g.State[pos.x][pos.y+1] != "E" {
			return
		}
	}

	for _, pos := range g.CurrentTetrominoes.positions {
		pos.y++
	}
}
