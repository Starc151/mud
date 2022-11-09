package player

import (
	conf "config"
	"fmt"
)

func PlayerStart() {
	data := conf.GetData("obj/users/players/playerData.json")
	startingLife := data.(map[string]interface{})
	fmt.Printf("HP: %.1f, MP: %.1f \n", startingLife["HP"], startingLife["MP"])
}
