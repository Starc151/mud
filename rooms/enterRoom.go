package rooms

import (
	"conf"
	"fmt"
)

func Description(description string, nameNPC string) {
	fmt.Printf("Вы находитесь в %s \n", description)
	if nameNPC != "" {
		fmt.Printf("Тут бродит %s\n", nameNPC)
	}
}
func EnterRoom(idRoom uint16) (map[string]uint16, string, string) {
	id := fmt.Sprint(idRoom)
	data := conf.GetData("rooms/json/" + id + ".json")
	roomData := data.(map[string]interface{})
	var npc map[string]interface{}
	var nameNPC string
	exits := map[string]uint16{
		"север":  conf.InterfaceFloatTouint16(roomData["north"]),
		"восток": conf.InterfaceFloatTouint16(roomData["east"]),
		"юг":     conf.InterfaceFloatTouint16(roomData["south"]),
		"запад":  conf.InterfaceFloatTouint16(roomData["west"]),
	}
	description := roomData["description"].(string)
	if roomData["npc"] != nil {
		npc = roomData["npc"].(map[string]interface{})
		nameNPC = npc["nameNPC"].(string)
	}
	return exits, description, nameNPC
}
