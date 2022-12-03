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

func Npc(idNpc uint16) NpcStruct {
	idNpcStr := fmt.Sprint(idNpc)
	npcData := conf.GetData("npc/" + idNpcStr + ".json")
	var Npc NpcStruct
	Npc.Name = npcData["name"].(string)
	Npc.Damage = conf.InterfaceFloatToUint16(npcData["damage"])
	Npc.Hp = conf.InterfaceFloatToUint16(npcData["hp"])
	Npc.Exp = conf.InterfaceFloatToUint16(npcData["exp"])
	fmt.Printf("Тут бродит %s\n", Npc.Name)
	return Npc
}
