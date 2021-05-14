package _interface

//消息接口
type Event interface {
	GetCode() int32
	SetCode(eventCode int32)
	GetSessionId() int32
	SetSessionId(sessionId int32)
	GetSession() interface{}
	SetSession(sess interface{})
	GetRoomId() int64
	SetRoomId(roomId int64)
	GetSeqId() int32
	SetSeqId(seqId int32)
	Encode() interface{}
	//FromMessage(obj interface{})  //构造消息
	Decode(obj interface{}) Event
}
