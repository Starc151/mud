package rooms

// import "fmt"

func ExitRoom(action string, exitsMap map[string]int) int {
	return exitsMap[action]
}
