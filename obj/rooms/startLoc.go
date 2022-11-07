package rooms

import (
	player "obj/users/players"
)

func StartLoc() {
	EnterRoom(13) // 13 starting room
	player.PlayerStart()
}
