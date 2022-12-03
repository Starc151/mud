package player

import (
	"fmt"
	"npc"
)

func attack(npc npc.NpcStruct, player PlayerStruct) PlayerStruct{
	fmt.Printf("Вы атаковали %s\n", npc.Name)
	for {
		npc.Hp -= player.Damage
		player.Hp -= npc.Damage
		fmt.Printf("❤ %d\n", player.Hp)
		if player.Hp <= 0 {
			fmt.Println("Упс...")
			Start()
		}
		if npc.Hp <=0 {
			fmt.Printf("Вы убили %s и получили %d опыта\n", npc.Name, npc.Exp)
			player.Exp += npc.Exp
			return player
		}
	}
}
