package player

import (
	"fmt"
	"obj/rooms"
)

func Action(idRoom, hp, mp int) {
	exits := rooms.EnterRoom(idRoom)
	fmt.Printf("HP: %d, MP: %d \n", hp, mp)
	var action string
	fmt.Println("Выши действия?")
	fmt.Scan(&action)
	if action == "выход" {
		return
	}
	exit := exits[action]
	if exit == 0 {
		fmt.Println("Так нельзя")
		Action(idRoom, hp, mp)
	}
	Action(exit, hp, mp)
}
