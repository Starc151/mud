package player

import "strings"

func instructions(action string) string {
	actionsMap := map[string]string{
		"север":  "north",
		"юг":     "south",
		"восток": "east",
		"запад":  "west",
		"выход":  "exitGame",
	}
	for kM, vM := range actionsMap {
		if strings.Contains(kM, action) {
			return vM
		}
	}
	return "err"
}
