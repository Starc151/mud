package player

import (
	"conf"
)

func PlayerStart() (hp, mp, attack uint16) {
	data := conf.GetData("players/playerData.json")
	startingLife := data.(map[string]interface{})
	hp = conf.InterfaceFloatTouint16(startingLife["HP"])
	mp = conf.InterfaceFloatTouint16(startingLife["MP"])
	attack = conf.InterfaceFloatTouint16(startingLife["attack"])
	return hp, mp, attack
}
