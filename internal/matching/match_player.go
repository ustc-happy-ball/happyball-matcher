package matching

import (
	"happyball-matcher/framework"
)

type MatchPlayer struct {
	id   int32
	sess *framework.BaseSession
}

func NewMatchPlayer(id int32, sess *framework.BaseSession) *MatchPlayer {
	return &MatchPlayer{id: id, sess: sess}
}
