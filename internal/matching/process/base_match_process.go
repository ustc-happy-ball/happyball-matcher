package process

import (
	"fmt"
	"happyball-matcher/configs"
	"happyball-matcher/dgs"
	"happyball-matcher/internal/event/info"
	"happyball-matcher/internal/event/response"
	"happyball-matcher/internal/matching/component"
	"happyball-matcher/internal/matching/rule"
	"strconv"
	"strings"
)

type BaseMatchProcess struct {
	rule component.MatchRule
}

func NewBaseMatchProcess() *BaseMatchProcess {
	return &BaseMatchProcess{
		rule: rule.NewPlayerCountMatchRule(1), //todo:config
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
	roomInfos, err := dgs.GetRoomList()
	if nil != err {
		fmt.Printf("[BaseMatchProcess]grpc拉取房间信息列表出错! err：%+v\n", err)
		return
	}
	//选取合适房间(房间最高分不超过胜利分数1/3)
	var targetRoomInfo *info.RoomInfo
	var targetRoomID int64
	for _, roomInfo := range roomInfos {
		if b.IsRoomAvailable(roomInfo) {
			targetRoomInfo = roomInfo
			targetRoomID = roomInfo.RoomId
			break
		}
	}
	if nil == targetRoomInfo {
		roomID, err := dgs.CreateRoom()
		if nil != err {
			fmt.Printf("[BaseMatchProcess]grpc创建房间信息出错! err：%+v\n", err)
			return
		}
		targetRoomID = roomID
	}
	//回包
	addrs := strings.Split(configs.DgsAddr, ":")
	if 2 != len(addrs) {
		fmt.Printf("[BaseMatchProcess]dgs地址解析出错! addr：%+v\n", configs.DgsAddr)
	}
	dgsIp := addrs[0]
	dgsPort,err := strconv.ParseInt(addrs[1], 10, 32)
	dgsAddr := info.NewConnectInfo(dgsIp, int32(dgsPort))
	resp := response.NewPlayerMatchingResponse(targetRoomID, dgsAddr)
	data :=	resp.ToGMessageBytes()
	for _, player := range players {
		player.Sess.Sess.Write(data)
	}
}

func (b *BaseMatchProcess) IsRoomAvailable(roomInfo *info.RoomInfo) bool {
	if nil == roomInfo {
		return false
	}
	if roomInfo.HighestScore <= 120 { //todo:config
		return true
	}
	return false
}
