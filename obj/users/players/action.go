package player

import (
	"fmt"
	"obj/rooms"
)

func Action(idRoom uint16, hp, mp, attack uint16) {
	exits, description, nameNPC := rooms.EnterRoom(idRoom)
	rooms.Description(description, nameNPC)
	rooms.ExitsRooms(exits)
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
				exits, description, nameNPC = rooms.EnterRoom(exits[action])
				rooms.Description(description, nameNPC)
				rooms.ExitsRooms(exits)
			}
		case "description":
			rooms.Description(description, nameNPC)
			rooms.ExitsRooms(exits)
		case "exitGame":
			fmt.Println("Вы вышли")
			return
		}
	}
}
