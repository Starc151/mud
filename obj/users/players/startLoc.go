package player

import (
	"fmt"
	"obj/rooms"
)

func Start() {
	PlayerStart()
	rooms.EnterRoom(13)
	var action int
	fmt.Scan(&action)
	Action(action)
}
