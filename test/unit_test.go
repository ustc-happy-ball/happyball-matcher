package test

import (
	"github.com/golang/protobuf/proto"
	"github.com/xtaci/kcp-go"
	pb "happyball-matcher/api/proto/pb"
	event2 "happyball-matcher/internal/event"
	"log"
	"time"
)

func SendDataToMatcher(data []byte) {
	if sess, err := kcp.DialWithOptions("127.0.0.1:8889", nil, 0, 0); err == nil {
		//sess调优
		sess.SetNoDelay(1, 10, 2, 1)
		sess.SetReadDeadline(time.Now().Add(time.Millisecond * time.Duration(2)))
		sess.SetACKNoDelay(true)
		//开启进入世界流程
		sess.Write(data)
		buf := make([]byte, 4096)
		num, _ := sess.Read(buf)
		if num > 0 {
			pbGMsg := &pb.MMessage{}
			proto.Unmarshal(buf, pbGMsg)
			msg := event2.MMessage{}
			m := msg.Decode(pbGMsg)
			log.Println(m)
			//buf清零
			for i := range buf {
				buf[i] = 0
			}
		}
	}
}
