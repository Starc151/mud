package npc

import (
	"conf"
	"fmt"
)

type NpcStruct struct {
	Name   string
	Hp     int
	Damage int
	Exp    int
}

func Npc(idNpc uint16) NpcStruct {
	idNpcStr := fmt.Sprint(idNpc)
	dataNpc := conf.GetData("npc/" + idNpcStr + ".json")
	var Npc NpcStruct
	Npc.Name = dataNpc["name"].(string)
	fmt.Printf("Тут бродит %s\n", Npc.Name)
	return Npc
}
