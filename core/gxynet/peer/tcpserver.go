package peer

import (
	"context"
	"net"
	"time"

	"github.com/zylikedream/galaxy/core/gxyconfig"
	"github.com/zylikedream/galaxy/core/gxynet/endpoint"
	"github.com/zylikedream/galaxy/core/gxynet/logger"
	"github.com/zylikedream/galaxy/core/gxyregister"
	"go.uber.org/zap"
)

type TcpServer struct {
	endpoint.CoreBundle
	listener net.Listener
	conf     *tcpServerConfig
}

type tcpServerConfig struct {
	Addr string `toml:"addr"`
}

func newTcpServer(c *gxyconfig.Configuration) (*TcpServer, error) {
	server := &TcpServer{}
	conf := &tcpServerConfig{}
	if err := c.UnmarshalKey(server.Type(), conf); err != nil {
		return nil, err
	}
	server.conf = conf
	if err := server.BindProc(c); err != nil {
		return nil, err
	}
	return server, nil
}

func (t *TcpServer) Init() error {
	return nil
}

func (t *TcpServer) Start(ctx context.Context, el endpoint.EventHandler) error {
	var err error
	t.listener, err = net.Listen("tcp", t.conf.Addr)
	if err != nil {
		return err
	}
	t.BindHandler(el)
	t.accept(ctx)
	return nil
}

func (t *TcpServer) accept(ctx context.Context) {
	for {
		con, err := t.listener.Accept()
		if err != nil {
			if nerr, ok := err.(net.Error); ok && nerr.Temporary() {
				logger.Nlog.Warn("tcpserver", zap.String("msg", "accept temporary error"))
				time.Sleep(time.Millisecond)
				continue
			}
			break
		}
		sess := endpoint.NewTcpEndPoint(con, t.CoreBundle)
		go func() {
			if err := sess.Start(ctx); err != nil {
				logger.Nlog.Warn("conn start faield", zap.Error(err))
			}
		}()
	}
}

func (t *TcpServer) Type() string {
	return PEER_TCP_SERVER
}

func (t *TcpServer) Build(c *gxyconfig.Configuration, args ...interface{}) (interface{}, error) {
	return newTcpServer(c)
}

func (t *TcpServer) Stop(ctx context.Context) {
}

func init() {
	gxyregister.Register((*TcpServer)(nil))
}
