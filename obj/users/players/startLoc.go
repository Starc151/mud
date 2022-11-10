package player

import "fmt"

func Start() {
	hp, mp := PlayerStart()
	fmt.Println(Action(13, hp, mp))
}
