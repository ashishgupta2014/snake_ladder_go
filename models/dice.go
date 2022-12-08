package models

import (
	"math/rand"
)

type Dice struct {
	no_of_dice int
	faces      int
}

func InitDice(no_of_dice int, faces int) *Dice {
	if no_of_dice == 0 {
		no_of_dice = 1
	}
	if faces == 0 {
		faces = 6
	}
	return &Dice{no_of_dice: no_of_dice, faces: faces}

}

func (dice Dice) RollDice() int {
	var score int
	var rolling_score int
	score = 0
	var possible_max_move = 3
	var max = 6
	var min = 1
	for possible_max_move > 0 {
		possible_max_move--
		var d = dice.no_of_dice
		rolling_score = 0
		for d > 0 {
			d--
			rolling_score += rand.Intn(max-min) + min
		}
		if rolling_score != dice.faces*dice.no_of_dice {
			return score + rolling_score
		}
		score += rolling_score
	}
	return 0
}
