package rooms

import "fmt"

func ExitsRooms(exit ExitsRoomStruct) {
	north := exitsLetter(exit.North, "С")
	east := exitsLetter(exit.East, "В")
	south := exitsLetter(exit.South, "Ю")
	west := exitsLetter(exit.West, "З")
	fmt.Printf("🏃: %s%s%s%s ", north, east, south, west)
}

func exitsLetter(exit uint16, letter string) string {
	if exit != 0 {
		return letter
	}
	return "_"
}
