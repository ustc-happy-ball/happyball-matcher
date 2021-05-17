package dgs

import (
	"context"
	"happyball-matcher/dgs/dgsGrpc"
	"happyball-matcher/internal/event/info"
)

var roomService dgsGrpc.RoomServiceClient

func getRoomService() (dgsGrpc.RoomServiceClient, error) {
	if roomService == nil {
		service, err := GetRoomGrpcService()
		if err != nil {
			return nil, err
		}
		roomService = service
		return service, nil
	}
	return roomService, nil
}

func GetRoomList() ([]*info.RoomInfo, error) {
	service, err := getRoomService()
	if err != nil {
		return nil, err
	}
	ctx := context.TODO()
	getRoomListReq := &dgsGrpc.GetRoomListRequest{}
	res, err := service.GetRoomList(ctx, getRoomListReq)
	if err != nil {
		return nil, err
	}
	roomMsgList := res.RoomList
	var roomList []*info.RoomInfo
	for _, roomMsg := range roomMsgList {
		room := &info.RoomInfo{}
		room = room.Decode(roomMsg).(*info.RoomInfo)
		roomList = append(roomList, room)
	}
	return roomList, nil
}

func CreateRoom() (int64, error) {
	service, err := getRoomService()
	if err != nil {
		return 0, err
	}
	ctx := context.TODO()
	createRoomReq := &dgsGrpc.CreateRoomRequest{}
	res, err := service.CreateRoom(ctx, createRoomReq)
	if err != nil {
		return 0, err
	}
	return res.RoomId, nil
}
