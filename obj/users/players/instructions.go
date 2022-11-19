package player

import "strings"

func instructions(action string) string {
	actionsMap := map[string]string{
		"север":       "север",
		"юг":          "юг",
		"восток":      "восток",
		"запад":       "запад",
		"выход":       "exitGame",
		"осмотреться": "description",
	}
	if len(action) <= 2 {
		return "short"
	}
	for kM, vM := range actionsMap {
		if strings.HasPrefix(kM, action) {
			return vM
		}
	}
	return "err"
}
