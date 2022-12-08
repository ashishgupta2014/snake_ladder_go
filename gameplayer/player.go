package gameplayer

import (
	"snake_ladder_game/models"
	"snake_ladder_game/display"
)

type Player struct {
	name     string
	position int
	status   bool
}

func InitPlayer(name string) *Player {
	return &Player{name: name, position: 0, status: false}

}


func (player *Player) Play(dice *models.Dice, board *models.Board) {
	var curr_pos = player.position
	var curr_score = dice.RollDice()
	var new_pos = curr_pos + curr_score
	new_pos = board.PositionOnBoard(new_pos)
	if new_pos != -1{
		player.position = new_pos
	}else{
		new_pos = curr_pos
	}
	display.PlayerMove(player.name, curr_score, curr_pos, new_pos)
}

func (player Player) GetStatus() bool{
	return player.status
}

func (player *Player) SetStatus(status bool){
	player.status = status
}

func (player *Player) GetPosition()int{
	return player.position
}

func (player *Player) GetName()string{
	return player.name
}