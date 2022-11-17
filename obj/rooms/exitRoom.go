package rooms

import "fmt"

func ExitsRooms(exits map[string]uint16) {
	north := exitsLetter(exits["север"], "С")
	east := exitsLetter(exits["восток"], "В")
	south := exitsLetter(exits["юг"], "Ю")
	west := exitsLetter(exits["запад"], "З")
	fmt.Printf("Выходы: %s%s%s%s\n", north, east, south, west)
}

func exitsLetter(exit uint16, letter string) string {
	if exit != 0 {
		return letter
	}
	return "_"
}
