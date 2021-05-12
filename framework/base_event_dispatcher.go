package framework
import (
	event "github.com/fwv/happyball-matcher/framework/event"
)
type BaseEventDispatcher struct {
	maxEventSize int32
	EventQueue   *event.EventRingQueue
}

func NewBaseEventDispatcher(maxEventSize int32) BaseEventDispatcher {
	return BaseEventDispatcher{
		maxEventSize: maxEventSize,
		EventQueue:   event.NewEventRingQueue(maxEventSize),
	}
}

func (b BaseEventDispatcher) FireEvent(e event.Event) {
	b.EventQueue.Push(e)
}

func (b BaseEventDispatcher) GetEventQueue() *event.EventRingQueue {
	return b.EventQueue
}

func (b BaseEventDispatcher) FireEventToSession(e event.Event, s event.Session) {

}

func (b BaseEventDispatcher) Close() {

}
