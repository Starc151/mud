package rooms

import (
	"conf"
	"fmt"
)

type ExitsRoomStruct struct {
	North uint16
	East  uint16
	South uint16
	West  uint16
}
type RoomStruct struct {
	Id          uint16
	Description string
	ExitsRoom   ExitsRoomStruct
	IdNpc       uint16
}

func EnterRoom(idRoom uint16) (room RoomStruct) {
	id := fmt.Sprint(idRoom)
	roomData := conf.GetData("rooms/json/" + id + ".json")
	room.ExitsRoom.North = conf.InterfaceFloatToUint16(roomData["north"])
	room.ExitsRoom.East = conf.InterfaceFloatToUint16(roomData["east"])
	room.ExitsRoom.South = conf.InterfaceFloatToUint16(roomData["south"])
	room.ExitsRoom.West = conf.InterfaceFloatToUint16(roomData["west"])
	room.Description = roomData["description"].(string)
	room.IdNpc = conf.InterfaceFloatToUint16(roomData["idNpc"])
	return room
}
