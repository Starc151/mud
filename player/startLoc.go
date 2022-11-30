package player

import (
	"fmt"
)

func Start() {
	const startLocation uint16 = 13
	player := PlayerStart()
	fmt.Println("Привет! Вы в игре!")
	actionSwitch(startLocation, player)
}
