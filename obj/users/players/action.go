package player

import (
	"fmt"
	"obj/rooms"
)

func Action(idRoom, hp, mp int) int {
	exits := rooms.EnterRoom(idRoom)
	action := ""
	fmt.Println("Ваши действия?")
	fmt.Scan(&action)
	action = instructions(action)

	switch action {
	case "err":
		fmt.Println("Так нельзя")
		return Action(idRoom, hp, mp)
	case "exitGame":
		fmt.Println("Вы вышли")
		return 0
	case "north", "south", "east", "west":
		exit := rooms.ExitRoom(action, exits)
		if exit == 0 {
			fmt.Println("Туда не пройти")
			return Action(idRoom, hp, mp)
		}
		return Action(exit, hp, mp)
	default:
		return Action(idRoom, hp, mp)
	}
}
