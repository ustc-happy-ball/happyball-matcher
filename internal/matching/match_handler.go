package matching

import (
	"fmt"
	pb "happyball-matcher/api/proto/pb"
	_interface "happyball-matcher/framework/interface"
	event2 "happyball-matcher/internal/event"
	"happyball-matcher/internal/event/request"
)

type MatchHandler struct {
	Matcher *Matcher
}

func NewMatchHandler() *MatchHandler {
	return &MatchHandler{}
}

func (m *MatchHandler) OnEvent(event _interface.Event) {
	if nil == event {
		return
	}
	msg := event.(*event2.MMessage)
	data := msg.Data
	switch data.GetCode() {
	case int32(pb.GAME_MSG_CODE_PLAYER_MATCHING_REQUEST):
		m.onPlayerMatching(data.(*request.PlayerMatchingRequest))

	}
}

func (m *MatchHandler) OnEventToSession(event _interface.Event, s _interface.Session) {
	panic("implement me")
}

func (m *MatchHandler) onPlayerMatching(req *request.PlayerMatchingRequest) {
	playerId := req.PlayerId
	if nil == req.Session {
		fmt.Errorf("[MatchHandler]处理英雄匹配时会话为空！")
	}
	//matchPlayer := NewMatchPlayer(playerId, req.Session)

	fmt.Println(playerId)
}
