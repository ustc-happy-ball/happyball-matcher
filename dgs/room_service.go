package dgs

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"happyball-matcher/dgs/dgsGrpc"
	"happyball-matcher/internal/event/info"
	"log"
	"time"
)

type ServiceConnection struct {
	conn        *grpc.ClientConn
	roomService dgsGrpc.RoomServiceClient
}

// InitConnection return a dgs gRPC client wrapper, it will block to initialize the connection
func InitConnection(address string) *ServiceConnection {
	ctx, _ := context.WithDeadline(context.TODO(), time.Now().Add(5*time.Second))
	conn, err := grpc.DialContext(ctx, address, grpc.WithInsecure())
	if err != nil {
		log.Fatalln("[grpc] 连接初始化失败，失败原因", err.Error())
	}

	return &ServiceConnection{
		conn:        conn,
		roomService: dgsGrpc.NewRoomServiceClient(conn),
	}
}

func (s *ServiceConnection) GetRoomGrpcService() (dgsGrpc.RoomServiceClient, error) {
	if s == nil {
		return nil, errors.New("没有初始化connection")
	}
	return s.roomService, nil
}

func (s *ServiceConnection) GetRoomList() ([]*info.RoomInfo, error) {
	ctx := context.TODO()
	getRoomListReq := &dgsGrpc.GetRoomListRequest{}
	res, err := s.roomService.GetRoomList(ctx, getRoomListReq)
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

func (s *ServiceConnection) CreateRoom() (int64, error) {
	ctx := context.TODO()
	createRoomReq := &dgsGrpc.CreateRoomRequest{}
	res, err := s.roomService.CreateRoom(ctx, createRoomReq)
	if err != nil {
		return 0, err
	}
	return res.RoomId, nil
}
