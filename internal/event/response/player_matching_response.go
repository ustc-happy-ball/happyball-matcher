package response

import (
	pb "happyball-matcher/api/proto/pb"
	"happyball-matcher/framework"
	_interface "happyball-matcher/framework/interface"
	"happyball-matcher/internal/event/info"
)

//message PlayerMatchingResponse {
//ConnectMsg dgsAddr = 1;//dgs服务器公网地址
//int32 RoomId = 2;
//}

type PlayerMatchingResponse struct {
	framework.BaseEvent //基础消息类作为父类
	RoomId              int32
	dgsAddr             *info.ConnectInfo
}

func NewPlayerMatchingResponse(roomId int32, dgsAddr *info.ConnectInfo) *PlayerMatchingResponse {
	return &PlayerMatchingResponse{RoomId: roomId, dgsAddr: dgsAddr}
}

func (p *PlayerMatchingResponse) Decode(obj interface{}) _interface.Event {
	pbMsg := obj.(*pb.Response).PlayerMatchingResponse
	p.RoomId = pbMsg.RoomId
	addrInfo := &info.ConnectInfo{}
	p.dgsAddr = addrInfo.Decode(pbMsg.DgsAddr).(*info.ConnectInfo)
	return p
}

func (p *PlayerMatchingResponse) Encode() interface{} {
	addrInfo := info.NewConnectInfo(p.dgsAddr.Ip, p.dgsAddr.Port)
	pbMsg := &pb.PlayerMatchingResponse{
		RoomId:  p.RoomId,
		DgsAddr: addrInfo.Encode().(*pb.ConnectMsg),
	}
	return pbMsg
}