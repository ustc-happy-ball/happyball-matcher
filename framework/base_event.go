package framework

type BaseEvent struct {
	SessionId int32
	Code      int32
	RoomId    int64
	SeqId     int32
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

