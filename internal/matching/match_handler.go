package matching

import (
	"fmt"
	pb "happyball-matcher/api/proto/pb"
	_interface "happyball-matcher/framework/interface"
	event2 "happyball-matcher/internal/event"
	"happyball-matcher/internal/event/request"
	"happyball-matcher/internal/matching/component"
	"log"
)

type MatchHandler struct {
	Matcher *Matcher
}

func NewMatchHandler(matcher *Matcher) *MatchHandler {
	return &MatchHandler{Matcher: matcher}
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
	matchPlayer := component.NewMatchPlayer(playerId, req.Session)
	if nil != m.Matcher {
		log.Printf("[MatchHandler]创建待匹配玩家，将该玩家加入匹配池！%+v\n", matchPlayer)
		err := m.Matcher.PutPlayerIntoMatchingPool(matchPlayer)
		if nil != err {
			log.Printf("[MatchHandler]玩家加入匹配池失败！error:%+v\n", err)
		}
		m.Matcher.process.StartMatching(m.Matcher.pool)
	}
}
