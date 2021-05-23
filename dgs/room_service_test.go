package dgs

import (
	"happyball-matcher/configs"
	"os"
	"testing"
)

func initProgram() {
	InitConnection("localhost:9000")
}

func TestCreateRoom(t *testing.T) {
	skipCI(t)
	initProgram()
	sa := InitConnection(configs.DgsAddr)
	roomId, err := sa.CreateRoom()
	if err != nil {
		t.Error("出错了:", err.Error())
	}
	t.Log("roomId为", roomId)
}

func TestGetRoomList(t *testing.T) {
	skipCI(t)
	initProgram()
	sa :=  InitConnection(configs.DgsAddr)
	roomList, err := sa.GetRoomList()
	if err != nil {
		t.Error("出错了:", err.Error())
	}
	t.Logf("roomList为%+v", roomList)
}



func skipCI(t *testing.T) {
	if os.Getenv("CI") != "" {
		t.Skip("Skipping testing in CI environment")
	}
}