package player

func Start() {
	const startLocation uint16 = 13
	player := PlayerStart()
	
	actionSwitch(startLocation, player)
}
