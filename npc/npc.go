package npc

import (
	"conf"
	"fmt"
)

type npcType struct {
	Name   string
	Hp     int
	Damage int
	Exp    int
}

func GetNpc(idNpc uint16) {
	dataNpc := conf.GetData("npc/skeleton.json")
	var npc npcType
	npc.Name = dataNpc["name"].(string)
	fmt.Printf("Тут бродит %s", npc.Name)
}
