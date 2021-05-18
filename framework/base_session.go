package framework

import (
	"errors"
	"github.com/xtaci/kcp-go"
	"sync"
	"time"
)

//基础会话类
type (
	BaseSession struct {
		Id                 int32           //唯一标识号，与player的ID相同
		SeqId              int32           //包序列号
		Sess               *kcp.UDPSession //kcp发送方
		Status             int32           //会话状态：建立、销毁
		Type               int32           //网络类型：TCP、UDP
		CreationTime       int64           //会话创建时间
		LastUpdateTime     int64           //上一次接收到消息的时间
		LastDisconnectTime int64           //会话上一次断开时间
		OfflineForever     bool            //超过30s没有发送消息即认为该player永久掉线了
		StatusMutex        sync.Mutex
	}
)

func NewBaseSession(id int32, s *kcp.UDPSession) *BaseSession {
	//kcp session调优
	s.SetNoDelay(1, 10, 2, 1)
	s.SetACKNoDelay(true)
	createTime := time.Now().UnixNano()
	baseSession := &BaseSession{
		Id:             id,
		Sess:           s,
		CreationTime:   createTime,
		LastUpdateTime: createTime,
	}
	return baseSession
}

//状态更新
func (c *BaseSession) SendMessage(buff []byte) error {
	if len(buff) > 1350 {
		return errors.New("发送的包太大了")
	}

	_, err := c.Sess.Write(buff)
	return err
}
