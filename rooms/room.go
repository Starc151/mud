package rooms

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type Room struct {
	Id          int
	Description string
	North       int
	East        int
	South       int
	West        int
}

func getDataRoom(idRoom int) []Room {
	file, err := os.Open("rooms/json/" + strconv.Itoa(idRoom) + ".json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	var roomData []Room
	errJ := json.Unmarshal(data, &roomData)
	if errJ != nil {
		panic(errJ)
	}
	return roomData
}

func EnterRoom(idRoom int) {
	roomData := getDataRoom(idRoom)
	fmt.Printf("Вы находитесь в %s.\n"+
		"Выходы:\n"+
		"HP: hp MP: mp",
		roomData[0].Description)
}
