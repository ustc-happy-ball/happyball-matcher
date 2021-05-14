package framework

import (
	"happyball-matcher/framework/interface"
	"sync"
)

type BaseEventManager struct {
	rw         *sync.RWMutex
	handlerMap map[int32]_interface.Handler
	eventMap   map[int32]_interface.Event
}

//static
var Manager = NewEventManager()

func NewEventManager() *BaseEventManager {
	return &BaseEventManager{
		rw:         &sync.RWMutex{},
		handlerMap: make(map[int32]_interface.Handler),
		eventMap:   make(map[int32]_interface.Event),
	}
}

func (e *BaseEventManager) RegisterEventToHandler(msgCode int32, event _interface.Event, handler _interface.Handler) {
	if nil != e.handlerMap[msgCode] {
		return
	}
	if nil != handler {
		e.rw.Lock()
		e.handlerMap[msgCode] = handler
		e.rw.Unlock()
	}
	if nil != e.eventMap[msgCode] {
		return
	}
	if nil != event {
		e.rw.Lock()
		e.eventMap[msgCode] = event
		e.rw.Unlock()
	}
}

func (e *BaseEventManager) FetchHandler(msgCode int32) _interface.Handler {
	return e.handlerMap[msgCode]
}

func (e *BaseEventManager) FetchEvent(msgCode int32) _interface.Event {
	return e.eventMap[msgCode]
}
