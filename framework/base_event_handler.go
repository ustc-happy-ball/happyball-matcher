package framework

import "happyball-matcher/framework/interface"

type BaseEventHandler struct {
}

func NewBaseEventHandler() *BaseEventHandler {
	return &BaseEventHandler{}
}

func (b *BaseEventHandler) OnEvent(e _interface.Event) {
	if nil == e {
		return
	}
	handler := Manager.FetchHandler(e.GetCode())
	if nil != handler {
		handler.OnEvent(e)
	}
}

func (b *BaseEventHandler) OnEventToSession(e _interface.Event, s _interface.Session) {
	if nil == e {
		return
	}
	handler := Manager.FetchHandler(e.GetCode())

	if nil != handler {
		handler.OnEvent(e)
	}
}
