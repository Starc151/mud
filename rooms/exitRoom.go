package rooms

import "fmt"

func ExitsRooms(exit ExitsRoomStruct) map[string]uint16 {
	north := exitsLetter(exit.North, "С")
	east := exitsLetter(exit.East, "В")
	south := exitsLetter(exit.South, "Ю")
	west := exitsLetter(exit.West, "З")
	fmt.Printf("🏃: %s%s%s%s\n", north, east, south, west)
	exitsRoomMap := map[string]uint16{
		"север":  exit.North,
		"восток": exit.East,
		"юг":     exit.South,
		"запад":  exit.West,
	}
	return exitsRoomMap
}


func exitsLetter(exit uint16, letter string) string {
	if exit != 0 {
		return letter
	}
	return "_"
}
