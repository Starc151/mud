package player

import (
	"fmt"
	"obj/rooms"
)

func Action(idRoom int) {
	rooms.EnterRoom(idRoom)
	var action int
	fmt.Scan(&action)
	if action < 1 {
		return
	}
	Action(action)
}
