package matching

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	pb "happyball-matcher/api/proto/pb"
	"happyball-matcher/configs"
	"happyball-matcher/framework"
	"happyball-matcher/framework/kcpnet"
	"happyball-matcher/internal/event"
	"happyball-matcher/internal/event/request"
	"happyball-matcher/internal/event/response"
	"log"
	"time"
)

type Matcher struct {
	pool       *MatchPool
	server     *kcpnet.KcpServer
	dispatcher *framework.BaseEventDispatcher
	handler    *framework.BaseEventHandler
}

func NewMatcher() *Matcher {
	s, err := kcpnet.NewKcpServer(configs.ServerAddr)
	if err != nil {
		return nil
	}
	m := &Matcher{
		server:     s,
		dispatcher: framework.NewBaseEventDispatcher(configs.MaxEventQueueSize),
		handler:    framework.NewBaseEventHandler(),
	}
	m.Init()
	return m
}

func (m *Matcher) Init() {
	log.Println("[Matcher]初始化系统组件！")

	matchHandler := NewMatchHandler()

	playerMatchingReq := request.PlayerMatchingRequest{}
	playerMatchingReq.SetCode(int32(pb.GAME_MSG_CODE_PLAYER_MATCHING_REQUEST))

	playerMatchingResp := response.PlayerMatchingResponse{}
	playerMatchingResp.SetCode(int32(pb.GAME_MSG_CODE_PLAYER_MATCHING_RESPONSE))

	framework.Manager.RegisterEventToHandler(int32(pb.GAME_MSG_CODE_PLAYER_MATCHING_REQUEST), &playerMatchingReq, matchHandler)
	framework.Manager.RegisterEventToHandler(int32(pb.GAME_MSG_CODE_PLAYER_MATCHING_RESPONSE), &playerMatchingResp, matchHandler)
}

func (m *Matcher) Serv() {
	log.Println("[Matcher]匹配器开始监听新连接！")
	buf := make([]byte, 1500)
	go m.HandleEventFromQueue()
	for {
		//select {
		//case :
		//	//
		//case <-time.After(time.Second * 1):
		//	//cpu保护
		//}
		conn, err := m.server.Listen.AcceptKCP()
		if err != nil {
			log.Println("[Matcher]kcp接收会话出错！")
		}
		session := framework.NewBaseSession(-1, conn)
		log.Printf("[GameRoomManager]监听到新连接！%v \n", session)
		err = session.Sess.SetReadDeadline(time.Now().Add(time.Millisecond * time.Duration(2)))
		if err != nil {
			log.Println("[Matcher]setDeadLine出错")
			continue
		}
		num, _ := session.Sess.Read(buf)
		if num == 0 {
			continue
		}
		pbMsg := &pb.MMessage{}
		err = proto.Unmarshal(buf[:num], pbMsg)
		if err != nil {
			log.Println("[Matcher]解析protobuf出错")
			log.Printf("%+v\n", err)
			continue
		}
		mmsg := event.MMessage{}
		mmsg.SetSession(session)
		msg := mmsg.Decode(pbMsg)
		//buf清零
		for i := range buf {
			buf[i] = 0
		}
		//放入消息队列中
		m.dispatcher.FireEvent(msg)
	}
}

func (m *Matcher) HandleEventFromQueue() {
	for {
		e, err := m.dispatcher.GetEventQueue().Pop()
		if nil == e { //todo
			continue
		}
		if nil != err {
			fmt.Println(err)
			continue
		}
		msg := e.(*event.MMessage)
		m.handler.OnEvent(msg)
	}
}
