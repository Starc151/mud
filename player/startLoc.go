package player

import (
	"fmt"
)

func Start() {
	const startLocation uint16 = 13
	hp, mp, damage := PlayerStart()
	fmt.Println("Привет! Вы в игре!")
	action(startLocation, hp, mp, damage)
}
