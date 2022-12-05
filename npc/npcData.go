package npc

import (
	"conf"
	"fmt"
)

type NpcStruct struct {
	Name   string
	Hp     uint16
	Damage uint16
	Exp    uint16
}

func NpcData(fileNpc string) NpcStruct {
	var Npc NpcStruct
	if fileNpc != "" {
		npcData := conf.GetData("npc/" + fileNpc + ".json")
		Npc.Name = npcData["name"].(string)
		Npc.Damage = conf.InterfaceFloatToUint16(npcData["damage"])
		Npc.Hp = conf.InterfaceFloatToUint16(npcData["hp"])
		Npc.Exp = conf.InterfaceFloatToUint16(npcData["exp"])
		fmt.Printf("Тут бродит %s\n", Npc.Name)
	}
	return Npc
}
