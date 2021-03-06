package framework

type BaseEvent struct {
	SessionId int32
	Code      int32
	RoomId    int64
	SeqId     int32
	Session   *BaseSession
}

func (e *BaseEvent) GetCode() int32 {
	return e.Code
}

func (e *BaseEvent) SetCode(code int32) {
	e.Code = code
}

func (e *BaseEvent) GetSessionId() int32 {
	return e.SessionId
}

func (e *BaseEvent) SetSessionId(sessionId int32) {
	e.SessionId = sessionId
}

func (e *BaseEvent) GetSession() interface{} {
	return e.Session
}

func (e *BaseEvent) SetSession(sess interface{}) {
	e.Session = sess.(*BaseSession)
}

func (e *BaseEvent) GetRoomId() int64 {
	return e.RoomId
}

func (e *BaseEvent) SetRoomId(roomId int64) {
	e.RoomId = roomId
}

func (e *BaseEvent) GetSeqId() int32 {
	return e.SeqId
}

func (e *BaseEvent) SetSeqId(seqId int32) {
	e.SeqId = seqId
}

//func NewBaseEvent(code int32) BaseEvent{
//	return BaseEvent{
//		Code: code,
//	}
//}

//func NewBaseEvent(code int32, sessionId int64) *BaseEvent{
//	return &BaseEvent{
//		Code: code,
//		SessionId: sessionId,
//	}
//}
