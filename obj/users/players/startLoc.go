package player

import (
	"fmt"
)

func Start() {
	hp, mp := PlayerStart()
	fmt.Println("Привет! Вы в игре!")
	Action(13, hp, mp)
}
