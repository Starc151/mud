package player

import "strings"

func instructions(action string) string {
	actionsMap := map[string]string{
		"север":  "север",
		"юг":     "юг",
		"восток": "восток",
		"запад":  "запад",
		"выход":  "exitGame",
	}
	for kM, vM := range actionsMap {
		if strings.Contains(kM, action) {
			return vM
		}
	}
	return "err"
}
