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

func GetNpc(idNpc uint16) {
	idNpcStr := fmt.Sprint(idNpc)
	dataNpc := conf.GetData("npc/" + idNpcStr + ".json")
	var Npc NpcStruct
	Npc.Name = dataNpc["name"].(string)
	fmt.Printf("Тут бродит %s", Npc.Name)
}
