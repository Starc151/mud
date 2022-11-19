package player

import (
	"bufio"
	"fmt"
	"obj/rooms"
	"os"
)

func actionScan() string {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	return in.Text()
}
func Action(idRoom uint16, hp, mp, attack uint16) {
	exits, description, nameNPC := rooms.EnterRoom(idRoom)
	rooms.Description(description, nameNPC)
	rooms.ExitsRooms(exits)
	action := ""
	for {
		if hp == 0 {
			fmt.Println("упс....")
			Start()
		}
		if hp <= 5 {

		}
		fmt.Printf("❤ %d 🧪 %d. Ваши действия?\n", hp, mp)
		action = instructions(actionScan())
		switch action {
		case "err":
			fmt.Println("Так нельзя")
		case "short":
			fmt.Println("Уточните...")
		case "север", "восток", "юг", "запад":
			fmt.Printf("Вы пошли на %s... ", action)
			if exits[action] != 0 {
				exits, description, nameNPC = rooms.EnterRoom(exits[action])
				rooms.Description(description, nameNPC)
				rooms.ExitsRooms(exits)
			} else {
				fmt.Println("Но туда нет прохода")
			}
		case "description":
			rooms.Description(description, nameNPC)
			rooms.ExitsRooms(exits)
		case "exitGame":
			fmt.Println("Вы точно хотите выйти?")
			var exitGame string
			fmt.Scan(&exitGame)
			if exitGame == "да" {
				fmt.Println("Вы вышли")
				return
			}
		}
	}
}
