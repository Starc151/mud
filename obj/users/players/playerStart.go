package player

import (
	conf "config"
)

func PlayerStart() (hp, mp, attack uint16) {
	data := conf.GetData("obj/users/players/playerData.json")
	startingLife := data.(map[string]interface{})
	hp = conf.InterfaceFloatTouint16(startingLife["HP"])
	mp = conf.InterfaceFloatTouint16(startingLife["MP"])
	attack = conf.InterfaceFloatTouint16(startingLife["attack"])
	return hp, mp, attack
}
