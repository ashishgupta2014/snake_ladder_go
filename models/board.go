package models

import (
	"snake_ladder_game/helper"
)

const SNAKE = "snake"
const LADDER = "ladder"

type Position struct {
	name  string
	point helper.Point
}

type Board struct {
	size int
	dict map[int]Position
}

func InitBoard(size int) *Board {
	return &Board{size: size, dict: make(map[int]Position)}
}

func (board *Board) SetSnake(points []helper.Point) {
	for _, point := range points {
		board.dict[point.GetX()] = Position{name: SNAKE, point: point}

	}

}

func (board *Board) SetLadder(points []helper.Point) {
	for _, point := range points {
		board.dict[point.GetX()] = Position{name: LADDER, point: point}
	}

}

func (board Board) PositionOnBoard(curr int) int {
	if curr > board.size {
		return -1
	}
	if _, ok := board.dict[curr]; ok {
		var position = board.dict[curr]
		curr = position.point.GetY()

	}
	return curr
}

func (board Board) IsEndPosition(pos int) bool {
	return board.size == pos
}
