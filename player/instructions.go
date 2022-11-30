package player

import "strings"

func instructionsString(str string) []string {
	strAction := strings.Split(str, " ")
	var action []string
	action = append(action, strAction...)
	return action
}
func lib(action []string) string {
	actionsList := map[string]string{
		"север":     "север",
		"юг":        "юг",
		"восток":    "восток",
		"запад":     "запад",
		"выход":     "exitGame",
		"осмотреть": "description",
		"атаковать": "attack",
		"помощь":    "help",
	}
	npcList := map[string]string{
		"скелет": "skeleton",
	}
	if len(action) == 1 {
		for k, v := range actionsList {
			if strings.HasPrefix(k, action[0]) {
				if v == "attack" {
					return "err"
				}
				return v
			}
		}
	} else {
		for k := range npcList {
			if strings.HasPrefix(k, action[1]) {
				return "attack"
			}
			return "err"
		}
	}
	return "err"
}
func instructions(actionPlayer string) string {
	action := instructionsString(actionPlayer)
	for _, v := range action {
		if len(v) <= 2 {
			return "err"
		}
	}
	return lib(action)
}
