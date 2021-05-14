package info

import (
	pb "happyball-matcher/api/proto/pb"
	"happyball-matcher/framework"
	_interface "happyball-matcher/framework/interface"
)

type ConnectInfo struct {
	framework.BaseEvent //基础消息类作为父类
	Ip                  string
	Port                int32
}

func NewConnectInfo(ip string, port int32) *ConnectInfo {
	return &ConnectInfo{
		Ip:   ip,
		Port: port,
	}
}

func (c *ConnectInfo) Decode(obj interface{}) _interface.Event {
	pbMsg := obj.(*pb.ConnectMsg)
	return &ConnectInfo{
		Ip:   pbMsg.Ip,
		Port: pbMsg.Port,
	}
}

func (c *ConnectInfo) Encode() interface{} {
	return &pb.ConnectMsg{
		Ip:   c.Ip,
		Port: c.Port,
	}
}
