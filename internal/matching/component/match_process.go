package component

import "happyball-matcher/internal/event/info"

type MatchProcess interface {
	StartMatching(pool *MatchPool)
	IsRoomAvailable(roomInfo *info.RoomInfo) bool
}
