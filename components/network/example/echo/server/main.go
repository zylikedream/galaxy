package main

import (
	"github.com/zylikedream/galaxy/components/glog"
	"github.com/zylikedream/galaxy/components/network"
	"github.com/zylikedream/galaxy/components/network/example/echo/proto"
	"github.com/zylikedream/galaxy/components/network/message"
	"github.com/zylikedream/galaxy/components/network/session"
	"go.uber.org/zap"
)

func main() {
	EchoServer()
}

type EchoEventHandler struct {
	session.BaseEventHandler
}

func (e *EchoEventHandler) OnOpen(sess session.Session) error {
	glog.Infof("session open, addr=%s", sess.Conn().RemoteAddr())
	return nil
}

func (e *EchoEventHandler) OnClose(sess session.Session) {
	glog.Infof("session close, addr=%s", sess.Conn().RemoteAddr())
}

func (e *EchoEventHandler) OnMessage(sess session.Session, msg *message.Message) error {
	switch m := msg.Msg.(type) {
	case *proto.EchoReq:
		glog.Infof("recv echo req:%v", m)
		sess.Send(&proto.EchoAck{
			Code: 0,
			Msg:  m.Msg,
		})
	}
	return nil
}

func EchoServer() {
	p, err := network.NewNetwork("config/config.toml")
	if err != nil {
		glog.Error("network", zap.Namespace("new failed"), zap.Error(err))
		return
	}
	if err := p.Start(&EchoEventHandler{}); err != nil {
		glog.Error("network", zap.Namespace("start failed"), zap.Error(err))
		return
	}

	done := make(chan struct{})
	<-done
}
