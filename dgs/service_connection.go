package dgs

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"happyball-matcher/dgs/dgsGrpc"
	"log"
	"time"
)

type ServiceConnection struct {
	conn        *grpc.ClientConn
	roomService dgsGrpc.RoomServiceClient
}

var serviceConnection *ServiceConnection

func InitConnection(address string) {
	serviceConnection = &ServiceConnection{}
	ctx, _ := context.WithDeadline(context.TODO(), time.Now().Add(5*time.Second))
	conn, err := grpc.DialContext(ctx, address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalln("[grpc] 连接初始化失败，失败原因", err.Error())
	}
	serviceConnection.conn = conn
	serviceConnection.roomService = dgsGrpc.NewRoomServiceClient(conn)
	log.Println("[grpc] grpc初始化成功")
}

func GetRoomGrpcService() (dgsGrpc.RoomServiceClient, error) {
	if serviceConnection == nil {
		return nil, errors.New("没有初始化connection")
	}
	return serviceConnection.roomService, nil
}
