package rooms

import (
	conf "config"
	"fmt"
	"strconv"
)

func exits(exit interface{}, numexit string) string {
	if exit != nil {
		return numexit
	}
	return "_"
}
func EnterRoom(idRoom int) (map[string]int, string) {
	id := strconv.Itoa(idRoom)
	data := conf.GetData("obj/rooms/json/" + id + ".json")
	roomData := data.(map[string]interface{})
	north := exits(roomData["north"], "С")
	east := exits(roomData["east"], "В")
	south := exits(roomData["south"], "Ю")
	west := exits(roomData["west"], "З")
	exits := map[string]int{
		"север":  conf.InterfaceFloatToInt(roomData["north"]),
		"восток": conf.InterfaceFloatToInt(roomData["east"]),
		"юг":     conf.InterfaceFloatToInt(roomData["south"]),
		"запад":  conf.InterfaceFloatToInt(roomData["west"]),
	}
	description := roomData["description"].(string)
	fmt.Printf("Вы находитесь в %s.\n", description)
	fmt.Printf("Выходы: %s%s%s%s\n", north, east, south, west)
	return exits, description
}
