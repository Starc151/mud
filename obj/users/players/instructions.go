package player

func instructions(action string) string {
	tmpStr := ""
	actionsMap := map[string]string{
		"север":  "north",
		"юг":     "south",
		"восток": "east",
		"запад":  "west",
		"выход":  "exitGame",
	}
	for kM, vM := range actionsMap {
		for _, v := range kM {
			tmpStr += string(v)
			if action == tmpStr {
				return vM

			}
		}
		tmpStr = ""
	}
	return "err"
}
