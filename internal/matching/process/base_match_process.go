package process

import (
	"happyball-matcher/internal/matching/component"
	"happyball-matcher/internal/matching/rule"
)

type BaseMatchProcess struct {
	rule component.MatchRule
}

func NewBaseMatchProcess() *BaseMatchProcess {
	return &BaseMatchProcess{
		rule: rule.NewPlayerCountMatchRule(2),
	}
}

func (b *BaseMatchProcess) StartMatching(pool *component.MatchPool) {
	players := make([]*component.MatchPlayer, 0)
	pool.PlayerMap.Range(func(key, value interface{}) bool {
		p := value.(*component.MatchPlayer)
		players = append(players, p)
		pool.PlayerMap.Delete(key)
		return true
	})
	res := b.rule.IsMatch(players)
	if !res { //未满足匹配条件则返回
		return
	}
	//拉取dgs房间信息,挑选一个合适的房间分配至玩家，如未找到则创建一个
}
