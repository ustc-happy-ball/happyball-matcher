package event

import (
	pb "happyball-matcher/api/proto/pb"
	"happyball-matcher/framework"
	"happyball-matcher/framework/interface"
	"log"
)

type MMessage struct {
	framework.BaseEvent
	MsgType     int32
	GameMsgCode int32
	SessionId   int32
	SeqId       int32
	SendTime    int64
	Data        _interface.Event
}

func (g *MMessage) Decode(obj interface{}) _interface.Event {
	pbMsg := obj.(*pb.MMessage)
	msg := &MMessage{
		MsgType: int32(pbMsg.MsgType),
		SeqId:   pbMsg.SeqId,
	}
	msg.SetCode(int32(pbMsg.MsgCode))
	msg.SetSessionId(pbMsg.SessionId)
	msg.SetSession(g.Session)
	msg.SetSeqId(pbMsg.SeqId)
	msg.SendTime = pbMsg.SendTime
	event := framework.Manager.FetchEvent(msg.GetCode())
	if nil == event {
		log.Printf("[MMessage]二级消息解包时未找到对应消息模板！请检查该类型消息是否在GameStarter中进行注册！pbMsg:%+v \n", pbMsg)
		return msg
	}
	if pb.MSG_TYPE_NOTIFY == pbMsg.MsgType {
		msg.Data = event.Decode(pbMsg.Notify)
	}
	if pb.MSG_TYPE_REQUEST == pbMsg.MsgType {
		msg.Data = event.Decode(pbMsg.Request)
	}
	if pb.MSG_TYPE_RESPONSE == pbMsg.MsgType {
		msg.Data = event.Decode(pbMsg.Response)
	}
	//传递会话id至二层协议中
	msg.Data.SetSessionId(pbMsg.SessionId)
	msg.Data.SetSeqId(pbMsg.SeqId)
	msg.Data.SetRoomId(g.RoomId)
	msg.Data.SetSession(g.Session)
	return msg
}

func (g *MMessage) Encode() interface{} {
	pbMsg := &pb.MMessage{}
	return pbMsg
}
