package info

import (
	"happyball-matcher/dgs/dgsGrpc"
	"happyball-matcher/framework"
	_interface "happyball-matcher/framework/interface"
)

type RoomInfo struct {
	framework.BaseEvent
	ID           int64
	Status       int32
	CreateTime   int64
	PlayerCount  int32
	HighestScore int32
}

func NewRoomInfo(ID int64, status int32, createTime int64, playerCount int32, highestScore int32) *RoomInfo {
	return &RoomInfo{ID: ID, Status: status, CreateTime: createTime, PlayerCount: playerCount, HighestScore: highestScore}
}

func (r RoomInfo) Encode() interface{} {
	return dgsGrpc.RoomMsg{
		Id:           r.ID,
		Status:       r.Status,
		CreateTime:   r.CreateTime,
		PlayerCount:  r.PlayerCount,
		HighestScore: r.HighestScore,
	}
}

func (r RoomInfo) Decode(obj interface{}) _interface.Event {
	pbMsg := obj.(*dgsGrpc.RoomMsg)
	return &RoomInfo{
		ID: pbMsg.Id,
		Status: pbMsg.Status,
		CreateTime: pbMsg.CreateTime,
		PlayerCount: r.PlayerCount,
		HighestScore: r.HighestScore,
	}
}
