package player

import (
	"bufio"
	"fmt"
	"helpGame"
	"npc"
	"os"
	"rooms"
)

func actionScan() string {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	return in.Text()
}
func actionSwitch(idRoom uint16, player PlayerStruct) {
	room := rooms.EnterRoom(idRoom)
	rooms.Description(room.Description)
	npcData := npc.NpcData(room.FileNpc)
	exitsRoomMap := rooms.ExitsRooms(room.ExitsRoom)
	var nextLvl uint16
	action := ""
	for {
		nextLvl = player.Lvl * 1100
		fmt.Printf("HP:%d EXP: %d|%d\n", player.Hp, player.Exp, nextLvl)
		fmt.Println("Ваши действия?")
		action = instructions(actionScan())
		switch action {
		case "err":
			fmt.Println("Уточните...")
		case "север", "восток", "юг", "запад":
			fmt.Printf("Вы пошли на %s... ", action)
			if exitsRoomMap[action] != 0 {
				room = rooms.EnterRoom(exitsRoomMap[action])
				rooms.Description(room.Description)
				npcData = npc.NpcData(room.FileNpc)
				exitsRoomMap = rooms.ExitsRooms(room.ExitsRoom)
			} else {
				fmt.Println("Но туда нет прохода")
			}
		case "description":
			rooms.Description(room.Description)
			npcData = npc.NpcData(room.FileNpc)
			rooms.ExitsRooms(room.ExitsRoom)
		case "exitGame":
			fmt.Println("Вы точно хотите выйти?")
			var exitGame string
			fmt.Scan(&exitGame)
			if exitGame == "да" {
				fmt.Println("Вы вышли")
				return
			}
		case "attack":
			if npcData.Name != "" {
				player = attack(npcData, player)
			} else {
				fmt.Println("Тут нет никого")
			}
		case "help":
			helpGame.HelpGame()
		}
	}
}
