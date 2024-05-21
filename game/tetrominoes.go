package game

import (
	"math/rand/v2"
)

var pieces = []TType{
	newTType('I', 20),
	newTType('Z', 10),
	newTType('J', 20),
	newTType('L', 20),
	newTType('O', 20),
	newTType('S', 10),
	newTType('T', 5), // I hate T shapes, or maybe it's just a skill issue
}

type Tetrominoes struct {
	positions []*position
	symbol    string
}

func NewTetrominoes() *Tetrominoes {
	piece := choice(pieces).symbol

	var positions []*position
	symbol := string([]byte{piece})
	pos := newPosition(-1, WIDTH/2)

	switch piece {
	case 'I':
		positions = []*position{pos, pos.up(), pos.up().up()}
	case 'Z':
		positions = []*position{pos, pos.up(), pos.up().left(), pos.right()}
	case 'J':
		cp := pos.up().up()
		positions = []*position{pos, pos.up(), cp, cp.left()}
	case 'L':
		cp := pos.left()
		positions = []*position{pos, cp, cp.up(), cp.up().up()}
	case 'O':
		positions = []*position{pos, pos.up(), pos.up().left(), pos.left()}
	case 'S':
		positions = []*position{pos, pos.up(), pos.up().right(), pos.left()}
	case 'T':
		cp := pos.up().up()
		positions = []*position{pos, pos.up(), cp, cp.left(), cp.right()}
	}

	return &Tetrominoes{positions, symbol}
}

type position struct {
	x, y int
}

func newPosition(x, y int) *position {
	return &position{x, y}
}

func (p position) up() *position {
	return &position{p.x - 1, p.y}
}

func (p position) left() *position {
	return &position{p.x, p.y - 1}
}

func (p position) right() *position {
	return &position{p.x, p.y + 1}
}

func (p position) down() *position {
	return &position{p.x + 1, p.y}
}

type TType struct {
	symbol byte
	weight int
}

func newTType(symbol byte, weight int) TType {
	return TType{symbol, weight}
}

func choice(array []TType) TType {
	sum := 0
	for _, e := range array {
		sum += e.weight
	}

	choice := rand.IntN(sum)

	sum = 0
	for _, e := range array {
		sum += e.weight

		if sum > choice {
			return e
		}
	}

	return array[0]
}
