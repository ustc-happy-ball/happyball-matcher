package kcpnet

import (
	"github.com/xtaci/kcp-go"
	"sync"
)

type KcpServer struct {
	mu     sync.Mutex
	addr   string
	Listen *kcp.Listener
	Sess   *kcp.UDPSession
}

// NewKcpServer return a *KcpServer
func NewKcpServer(addr string) (s *KcpServer, err error) {
	ts := new(KcpServer)
	ts.addr = addr
	ts.Listen, err = kcp.ListenWithOptions(addr, nil, 0, 0)
	if err != nil {
		return nil, err
	}
	return ts, err
}
