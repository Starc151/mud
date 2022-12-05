package player

import "fmt"

func lvlUp(player PlayerStruct) PlayerStruct {
	player.Lvl = (player.Exp / 1100) + 1
	player.Hp += player.Hp / 10
	player.Damage += player.Damage/10 + 1
	fmt.Printf("Теперь у вас %d уровень, %d HP и %d атаки\n",
		player.Lvl, player.Hp, player.Damage)
	return player
}
