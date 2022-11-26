package player

import (
	"fmt"
)

func Start() {
	const startLocation uint16 = 13
	hp, mp, attack := PlayerStart()
	fmt.Println("Привет! Вы в игре!")
	action(startLocation, hp, mp, attack)
}
