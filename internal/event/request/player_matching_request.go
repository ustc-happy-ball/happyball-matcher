package request

import (
	"github.com/golang/protobuf/proto"
	pb "happyball-matcher/api/proto/pb"
	"happyball-matcher/framework"
	_interface "happyball-matcher/framework/interface"
)

//message PlayerMatchingRequest {
//int32 PlayerId = 1; //玩家ID
//}

type PlayerMatchingRequest struct {
	framework.BaseEvent //基础消息类作为父类
	PlayerId            int32
}

func NewPlayerMatchingRequest(playerId int32) *PlayerMatchingRequest {
	return &PlayerMatchingRequest{PlayerId: playerId}
}

func (p *PlayerMatchingRequest) Decode(obj interface{}) _interface.Event {
	pbMsg := obj.(*pb.Request).PlayerMatchingRequest
	p.PlayerId = pbMsg.PlayerId
	return p
}

func (p *PlayerMatchingRequest) Encode() interface{} {
	return &pb.PlayerMatchingRequest{
		PlayerId: p.PlayerId,
	}
}

func (p *PlayerMatchingRequest) ToGMessageBytes() []byte {
	req := &pb.Request{
		PlayerMatchingRequest: p.Encode().(*pb.PlayerMatchingRequest),
	}
	msg := pb.MMessage{
		MsgType: pb.MSG_TYPE_REQUEST,
		MsgCode: pb.GAME_MSG_CODE_PLAYER_MATCHING_REQUEST,
		Request: req,
		SeqId:   -1,
	}
	out, _ := proto.Marshal(&msg)
	return out
}
