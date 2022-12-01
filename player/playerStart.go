package player

import (
	"conf"
)

type PlayerStruct struct {
	Name   string
	Hp     uint16
	Mp     uint16
	Damage uint16
	Exp    uint16
	Lvl    uint16
}

func PlayerStart() PlayerStruct {
	startingLife := conf.GetData("player/playerData.json")
	var player PlayerStruct
	player.Hp = conf.InterfaceFloatToUint16(startingLife["HP"])
	player.Mp = conf.InterfaceFloatToUint16(startingLife["MP"])
	player.Damage = conf.InterfaceFloatToUint16(startingLife["damage"])
	player.Exp = conf.InterfaceFloatToUint16(startingLife["exp"])
	player.Lvl = conf.InterfaceFloatToUint16(startingLife["lvl"])

	return player
}
