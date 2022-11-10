package rooms

import (
	conf "config"
	"fmt"
	"strconv"
)

func EnterRoom(idRoom int) map[string]int {
	id := strconv.Itoa(idRoom)
	data := conf.GetData("obj/rooms/json/" + id + ".json")
	roomData := data.(map[string]interface{})
	north, east, south, west := "_", "_", "_", "_"
	if roomData["north"] != nil {
		north = "С"
	}
	if roomData["east"] != nil {
		east = "В"
	}
	if roomData["south"] != nil {
		south = "Ю"
	}
	if roomData["west"] != nil {
		west = "З"
	}
	exits := map[string]int{
		"с": conf.InterfaceFloatToInt(roomData["north"]),
		"в": conf.InterfaceFloatToInt(roomData["east"]),
		"ю": conf.InterfaceFloatToInt(roomData["south"]),
		"з": conf.InterfaceFloatToInt(roomData["west"]),
	}
	fmt.Printf("Вы находитесь в %s.\n", roomData["description"])
	fmt.Printf("Выходы: %s%s%s%s\n", north, east, south, west)
	return exits
}