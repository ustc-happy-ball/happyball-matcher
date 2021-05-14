package matching

import "sync"

type MatchPool struct {
	PlayerMap sync.Map //待匹配玩家集合
}
