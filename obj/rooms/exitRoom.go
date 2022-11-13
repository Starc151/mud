package rooms

import "fmt"

func ExitRoom(action string, exitsMap map[string]int) int {
	fmt.Printf("Вы пошли на %s... ", action)
	return exitsMap[action]
}
