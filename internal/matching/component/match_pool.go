package component

import (
	"sync"
)

type MatchPool struct {
	PlayerMap sync.Map //待匹配玩家集合
}

func NewMatchPool() *MatchPool {
	return &MatchPool{}
}

func (p *MatchPool) AddPlayer(player *MatchPlayer) {
	p.PlayerMap.Store(player.Id, player)
}

func (p *MatchPool) QueryPlayer(playerId int32) *MatchPlayer {
	res, _ := p.PlayerMap.Load(playerId)
	return res.(*MatchPlayer)
}
