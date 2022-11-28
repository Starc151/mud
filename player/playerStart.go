package player

import (
	"conf"
)

func PlayerStart() (hp, mp, damage uint16) {
	data := conf.GetData("player/playerData.json")
	startingLife := data.(map[string]interface{})
	hp = conf.InterfaceFloatTouint16(startingLife["HP"])
	mp = conf.InterfaceFloatTouint16(startingLife["MP"])
	damage = conf.InterfaceFloatTouint16(startingLife["attack"])
	return hp, mp, damage
}
