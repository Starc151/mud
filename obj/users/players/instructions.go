package player

import (
	"fmt"
	"strings"
)

func instructionsQuantity(str string) []string {
	strAction := strings.Split(str, " ")
	var action []string
	for _, v := range strAction {
		action = append(action, v)
	}
	return action
}
func lib(action string) string {
	actionsMap := map[string]string{
		"север":       "север",
		"юг":          "юг",
		"восток":      "восток",
		"запад":       "запад",
		"выход":       "exitGame",
		"осмотреться": "description",
		"атаковать":   "attack",
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
func instructions(actionP string) string {
	action := instructionsQuantity(actionP)
	if len(action) == 1 {
		return lib(action[0])
	}
	if action[0] == "attack" && len(action) == 1 {
		lib(action[0]
		fmt.Print("Rjuj")
	}
	return "err"
}
