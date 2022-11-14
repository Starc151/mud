package player

import (
	conf "config"
)

func PlayerStart() (int, int, int) {
	data := conf.GetData("obj/users/players/playerData.json")
	startingLife := data.(map[string]interface{})
	hp := conf.InterfaceFloatToInt(startingLife["HP"])
	mp := conf.InterfaceFloatToInt(startingLife["MP"])
	attack := conf.InterfaceFloatToInt(startingLife["attack"])
	return hp, mp, attack
}
