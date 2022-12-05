package player

import (
	"fmt"
	"npc"
	"time"
)

func attack(npc npc.NpcStruct, player PlayerStruct) PlayerStruct {
	fmt.Printf("Вы вероломно атаковали %s\n", npc.Name)
	// nextLvl := player.Lvl * 1100
	for {
		time.Sleep(1000 * time.Millisecond)
		if npc.Hp <= player.Damage {
			npc.Hp = 0
			player.Exp += npc.Exp
			fmt.Printf("Вы нанесли смертельный удао и зверски убили беззащитного %sа\n", npc.Name)
			fmt.Printf("Вы получили %d опыта\n", npc.Exp)
			return lvlUp(player)
		}
		npc.Hp -= player.Damage
		fmt.Printf("Вы нанесли %d урона, HP %sа: %d\n", player.Damage, npc.Name, npc.Hp)
		time.Sleep(1000 * time.Millisecond)
		if player.Hp <= npc.Damage {
			fmt.Println("Упс...")
			Start()
		}
		player.Hp -= npc.Damage
		fmt.Printf("Вам нанесено %d урона, HP: %d\n", npc.Damage, player.Hp)
	}
}
