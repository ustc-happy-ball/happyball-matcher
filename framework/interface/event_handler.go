package _interface

type Handler interface {
	//处理消息
	OnEvent(event Event)
	OnEventToSession(event Event, s Session)
}
