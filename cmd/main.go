package main

import (
	"rooms"
	player "users/players"
)

type Player struct {
	Name string
	HP   int
	MP   int
}

func main() {
	rooms.EnterRoom(13)
	player.PlayerLife()
}
