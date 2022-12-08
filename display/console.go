package display

import (
	"fmt"
)

func PlayerMove(player string, dice_face int, initial_pos int, final_pos int){
	
	if initial_pos != final_pos{
		fmt.Println(player, " rolled a ", dice_face, " and moved from ", initial_pos, " to ", final_pos)
	}else{
		fmt.Println(player, " rolled a ", dice_face, " and but can't move from ", initial_pos)
		
	}

}