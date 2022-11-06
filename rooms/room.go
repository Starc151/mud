package rooms

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type Room struct {
	Id, North, East, South, West int
	Description                  string
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
	north, east, south, west := "", "", "", ""
	if roomData[0].North != 0 {
		north = "Север"
	}
	if roomData[0].East != 0 {
		east = "Восток"
	}
	if roomData[0].South != 0 {
		south = "Юг"
	}
	if roomData[0].West != 0 {
		west = "Запад"
	}

	fmt.Printf("Вы находитесь в %s.\n", roomData[0].Description)
	fmt.Printf("Выходы: %s %s %s %s\n", north, east, south, west)
}
