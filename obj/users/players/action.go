package player

import (
	"fmt"
	"obj/rooms"
)

func Actionn() {
	action := ""
	exits := rooms.EnterRoom(13) // 13 starting room
	PlayerStart()
	fmt.Print("Ваши действия? ")
	fmt.Scanln(&action)
	exit := exits[action]
	rooms.EnterRoom(exit)

}
