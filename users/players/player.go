package player

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Player struct {
	Name string
	HP   int
	MP   int
}

func getPlayerData() []Player {
	file, err := os.Open("users/players/playerData.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	var result []Player
	errJ := json.Unmarshal(data, &result)
	if errJ != nil {
		panic(errJ)
	}
	return result
}
func PlayerLife() {
	life := getPlayerData()
	fmt.Printf("HP: %d MP: %d", life[0].HP, life[0].MP)
}
