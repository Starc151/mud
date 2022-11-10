package player

import (
	"fmt"
	"obj/rooms"
)

func Action(idRoom, hp, mp int) string {
	exits := rooms.EnterRoom(idRoom)
	hp -= 1
	fmt.Printf("HP: %d, MP: %d \n", hp, mp)
	var action string
	fmt.Println("Выши действия?")
	fmt.Scan(&action)
	if action == "выход" {
		return "Вы вышли"
	}

	exit := exits[action]
	if exit == 0 {
		fmt.Println("Так нельзя")
		return Action(idRoom, hp, mp)
	}
	return Action(exit, hp, mp)
}
