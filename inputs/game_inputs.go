package inputs

import (
	"strconv"
	"strings"
)

type RawInputs struct{
	board_size int
	no_of_dice int
	no_of_snakes int
	snakes [][2]int
	no_of_ladders int
	ladders [][2]int
	no_of_players int
	players []string
}

func InitInput(content string)*RawInputs{
	arr := strings.Split(string(content), "\n")
	j := 0
	board_size, _ := strconv.Atoi(arr[j])
	j++
	no_of_dice, _ := strconv.Atoi(arr[j])
	j++
	no_of_snakes, _ := strconv.Atoi(arr[j])
	j++
	var snakes = make([][2]int, no_of_snakes)

	for i:=0; i<no_of_snakes; i++{
		snake := strings.Split(arr[j], " ")
		snakes[i][0], _ = strconv.Atoi(snake[0])
		snakes[i][1], _ = strconv.Atoi(snake[1])
		j++
	}
	no_of_ladders, _ := strconv.Atoi(arr[j])
	j++
	var ladders = make([][2]int, no_of_ladders)
	for i:=0; i<no_of_ladders; i++{
		ladder := strings.Split(arr[j], " ")
		ladders[i][0], _ = strconv.Atoi(ladder[0])
		ladders[i][1], _ = strconv.Atoi(ladder[1])
		j++
	}
	no_of_players, _ := strconv.Atoi(arr[j])
	j++
	var players = make([]string, no_of_players)
	for i:=0; i<no_of_players; i++{
		players[i] = arr[j]
		j++
	}
	raw_input := RawInputs{
		board_size: board_size,
		no_of_dice: no_of_dice,
		no_of_snakes: no_of_snakes,
		snakes: make([][2]int, no_of_snakes),
		no_of_ladders: no_of_ladders,
		ladders: make([][2]int, no_of_ladders),
		no_of_players: no_of_players,
		players: make([]string, no_of_players),
	}
	raw_input.snakes = snakes
	raw_input.ladders = ladders
	raw_input.players = players
	return &raw_input
	
}

func (raw_input *RawInputs) GetSizeBoard()int{
	return raw_input.board_size
}

func (raw_input *RawInputs) NoOfDice()int{
	return raw_input.no_of_dice
}

func (raw_input *RawInputs) NoOfPlayers()int{
	return raw_input.no_of_players
}

func (raw_input *RawInputs) Players(index int)string{
	return raw_input.players[index]
}


func (raw_input *RawInputs) NoOfSnakes()int{
	return raw_input.no_of_snakes
}

func (raw_input *RawInputs) GetSnakes()[][2]int{
	return raw_input.snakes
}

func (raw_input *RawInputs) NoOfLadders()int{
	return raw_input.no_of_ladders
}

func (raw_input *RawInputs) GetLadders()[][2]int{
	return raw_input.ladders
}