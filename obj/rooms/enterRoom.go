package rooms

import (
	conf "config"
	"fmt"
	"strconv"
)

func Description(description string, nameNPC string) {
	if nameNPC == "" {
		fmt.Printf("Вы находитесь в %s \n", description)
	} else {
		fmt.Printf("Вы находитесь в %s \nТут бродит %s\n", description, nameNPC)
	}
}
func EnterRoom(idRoom int) (map[string]int, string, string) {
	id := strconv.Itoa(idRoom)
	data := conf.GetData("obj/rooms/json/" + id + ".json")
	roomData := data.(map[string]interface{})
	var npc map[string]interface{}
	var nameNPC string
	exits := map[string]int{
		"север":  conf.InterfaceFloatToInt(roomData["north"]),
		"восток": conf.InterfaceFloatToInt(roomData["east"]),
		"юг":     conf.InterfaceFloatToInt(roomData["south"]),
		"запад":  conf.InterfaceFloatToInt(roomData["west"]),
	}
	description := roomData["description"].(string)
	if roomData["npc"] != nil {
		npc = roomData["npc"].(map[string]interface{})
		nameNPC = npc["nameNPC"].(string)
	}
	return exits, description, nameNPC
}
