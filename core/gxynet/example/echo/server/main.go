package main

import (
	"context"

	"github.com/zylikedream/galaxy/core/gxyconfig"
	"github.com/zylikedream/galaxy/core/gxylog"
	"github.com/zylikedream/galaxy/core/gxynet"
	"github.com/zylikedream/galaxy/core/gxynet/example/echo/proto"
	"github.com/zylikedream/galaxy/core/gxynet/message"
	"github.com/zylikedream/galaxy/core/gxynet/session"
	"go.uber.org/zap"
)

func main() {
	EchoServer()
}

type EchoEventHandler struct {
	session.BaseEventHandler
}

func (e *EchoEventHandler) OnOpen(ctx context.Context, sess session.Session) error {
	gxylog.Infof("session open, addr=%s", sess.Conn().RemoteAddr())
	return nil
}

func (e *EchoEventHandler) OnClose(ctx context.Context, sess session.Session) {
	gxylog.Infof("session close, addr=%s", sess.Conn().RemoteAddr())
}

func (e *EchoEventHandler) OnMessage(ctx context.Context, sess session.Session, msg *message.Message) error {
	switch m := msg.Msg.(type) {
	case *proto.EchoReq:
		gxylog.Infof("recv echo req:%v", m)
		sess.Send(&proto.EchoAck{
			Code: 0,
			Msg:  m.Msg,
		})
	}
	return nil
}

func EchoServer() {
	p, err := gxynet.NewNetwork(gxyconfig.New("config/gxynet.toml"))
	if err != nil {
		gxylog.Error("gxynet", zap.Namespace("new failed"), zap.Error(err))
		return
	}
	if err := p.Start(context.Background(), &EchoEventHandler{}); err != nil {
		gxylog.Error("gxynet", zap.Namespace("start failed"), zap.Error(err))
		return
	}
}