package framework

import "happyball-matcher/framework/interface"

type BaseEventDispatcher struct {
	maxEventSize int32
	EventQueue   *EventRingQueue
}

func NewBaseEventDispatcher(maxEventSize int32) *BaseEventDispatcher {
	return &BaseEventDispatcher{
		maxEventSize: maxEventSize,
		EventQueue:   NewEventRingQueue(maxEventSize),
	}
}

func (b *BaseEventDispatcher) FireEvent(e _interface.Event) {
	b.EventQueue.Push(e)
}

func (b *BaseEventDispatcher) GetEventQueue() *EventRingQueue {
	return b.EventQueue
}

func (b *BaseEventDispatcher) FireEventToSession(e _interface.Event, s _interface.Session) {

}

func (b *BaseEventDispatcher) Close() {

}
