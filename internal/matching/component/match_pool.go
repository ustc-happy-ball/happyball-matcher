package component

import (
	"errors"
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

func (p *MatchPool) QueryPlayer(playerId int32) (*MatchPlayer, error){
	res, _ := p.PlayerMap.Load(playerId)
	if nil == res {
		return nil, errors.New("玩家不存在！")
	}
	return res.(*MatchPlayer), nil
}
