package player

import (
	"fmt"
	"npc"
	"time"
)

func attack(npc npc.NpcStruct, player PlayerStruct) PlayerStruct {
	fmt.Printf("Вы атаковали %s\n", npc.Name)
	nextLvl := player.Lvl * 1100
	for {
		time.Sleep(1000 * time.Millisecond)
		npc.Hp -= player.Damage
		player.Hp -= npc.Damage
		fmt.Printf("❤ %d  |  ❤ %d\n", player.Hp, npc.Hp)
		if player.Hp <= 0 {
			player.Exp -= player.Exp / 10
			fmt.Println("Упс...")
			Start()
		}
		if npc.Hp <= 0 {
			fmt.Printf("Вы зверски убили беззащитный %s\n", npc.Name)
			fmt.Printf("и получили %d опыта\n", npc.Exp)
			player.Exp += npc.Exp
			if player.Exp >= nextLvl {
				player.Lvl = (player.Exp / 1100) + 1
				fmt.Printf("Ваш уроыень теперь %d\n", player.Lvl)
			}
			return player
		}
	}
}
