package component

import (
	"happyball-matcher/framework"
)

type MatchPlayer struct {
	Id   int32
	Sess *framework.BaseSession
}

func NewMatchPlayer(id int32, sess *framework.BaseSession) *MatchPlayer {
	return &MatchPlayer{Id: id, Sess: sess}
}
