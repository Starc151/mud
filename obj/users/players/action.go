package player

import (
	"fmt"
	"obj/rooms"
)

func Action(idRoom, hp, mp int) {
	exits := rooms.EnterRoom(idRoom)
	action := ""
	for {
		fmt.Println("Ваши действия?")
		fmt.Scan(&action)
		action = instructions(action)
		switch action {
		case "err":
			fmt.Println("Так нельзя")
		case "север", "восток", "юг", "запад":
			fmt.Printf("Вы пошли на %s... ", action)
			if exits[action] == 0 {
				fmt.Println("но туда нет прохода")
			} else {
				exits = rooms.EnterRoom(exits[action])
			}
		default:
			fmt.Println("Вы вышли")
			return
		}
	}
}
