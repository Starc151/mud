package player

import (
	"bufio"
	"fmt"
	"helpGame"
	"os"
	"rooms"
)

func actionScan() string {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	return in.Text()
}
func actionSwitch(idRoom uint16, player paramPlayer) {
	exits, description, nameNPC := rooms.EnterRoom(idRoom)
	rooms.Description(description, nameNPC)
	rooms.ExitsRooms(exits)
	action := ""
	for {
		if player.Hp == 0 {
			fmt.Println("упс....")
			Start()
		}
		fmt.Printf("❤ %d 🧪 %d. Ваши действия?\n", player.Hp, player.Mp)
		action = instructions(actionScan())
		switch action {
		case "err":
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
		case "attack":
			if nameNPC != "" {
				fmt.Printf("Вы атаковали %s\n", nameNPC)
			} else {
				fmt.Println("Тут нет никого")
			}
		case "help":
			helpGame.HelpGame()
		}
	}
}
