package dgs

import (
	"testing"
)

func initProgram() {
	InitConnection("localhost:9000")
}

func TestCreateRoom(t *testing.T) {
	initProgram()
	roomId, err := CreateRoom()
	if err != nil {
		t.Error("出错了:", err.Error())
	}
	t.Log("roomId为", roomId)
}

func TestGetRoomList(t *testing.T) {
	initProgram()
	roomList, err := GetRoomList()
	if err != nil {
		t.Error("出错了:", err.Error())
	}
	t.Logf("roomList为%+v", roomList)
}
