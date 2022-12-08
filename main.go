package main

import (
	"fmt"
	"snake_ladder_game/gameplayer"
	"snake_ladder_game/helper"
	"snake_ladder_game/models"
)

var board_size int

var no_of_dices int

var no_of_players int

var players []*gameplayer.Player

var no_of_snakes int

var snakes []helper.Point

var no_of_ladders int

var ladders []helper.Point

func init() {
	game := helper.InputReader("public/input_1.txt")
	fmt.Println("Snake and Ladder Game")
	board_size = game.GetSizeBoard()
	fmt.Println("Board Size: ", board_size)

	no_of_dices = game.NoOfDice()
	fmt.Println("No of Dices: ", no_of_dices)

	no_of_players = game.NoOfPlayers()
	fmt.Println("No of players: ", no_of_players)

	players = make([]*gameplayer.Player, no_of_players)

	for i := 0; i < no_of_players; i++ {
		players[i] = gameplayer.InitPlayer(game.Players(i))
	}
	// players[0] = gameplayer.InitPlayer("Ashu")
	// players[1] = gameplayer.InitPlayer("Rani")
	// players[2] = gameplayer.InitPlayer("Avi")

	no_of_snakes = game.NoOfSnakes()

	fmt.Println("No of Snakes in Game: ", no_of_snakes)
	//snakes = make([]helper.Point, no_of_snakes)
	var snake_input_raw_array = [][2]int{{62, 5}, {33, 6}, {49, 9}, {88, 16}, {41, 20}, {56, 53}, {98, 64}, {93, 73}, {95, 75}}
	//var snake_input_raw_array = game.GetSnakes()
	snakes = helper.BuildPointerArray(no_of_snakes, snake_input_raw_array)
	// fmt.Println(snakes)

	no_of_ladders = game.NoOfLadders()
	fmt.Println("No of Ladder in Game: ", no_of_ladders)
	ladders = make([]helper.Point, no_of_ladders)
	//var ladder_input_raw_array = [][2]int{{2, 37}, {27, 46}, {10, 32}, {51, 68}, {61, 79}, {65, 84}, {71, 91}, {81, 100}}
	var ladder_input_raw_array = game.GetLadders()
	ladders = helper.BuildPointerArray(no_of_ladders, ladder_input_raw_array)
	//fmt.Println(ladders)

}

func main() {
	var board = models.InitBoard(board_size)
	var dice = models.InitDice(1, 6)
	board.SetSnake(snakes)
	board.SetLadder(ladders)

	play_game(board, players, dice)
}

func play_game(board *models.Board, players []*gameplayer.Player, dice *models.Dice) {
	var winner = make([]string, no_of_players-1)
	var w = 0
	var active_player = no_of_players
	for active_player > 1 {
		for _, p := range players {
			if p.GetStatus() == false {
				p.Play(dice, board)
				if board.IsEndPosition(p.GetPosition()) {
					p.SetStatus(true)
					active_player--
					winner[w] = p.GetName()
					w++
				}
			}

		}

	}

	fmt.Println("-----Game Over-----")
	for i, w := range winner {
		fmt.Println(i+1, "Winner is", w)
	}
}
