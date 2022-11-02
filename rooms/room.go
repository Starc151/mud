package rooms

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

func GetDataRoom(idRoom int) {
	file, err := os.Open("src/rooms/json/" + strconv.Itoa(idRoom) + ".json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	data := make([]byte, 64)
	for {
		n, err := file.Read(data)
		if err == io.EOF {
			break
		}
		fmt.Print(string(data[:n]))
	}

}

func EnterRoom(description string, i int) {
	fmt.Printf("Вы находитесь в %s комнате.\nВыходы:\nHP: hp MP: mp", description)
}
