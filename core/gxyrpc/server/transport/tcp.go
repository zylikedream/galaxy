package transport

import (
	"github.com/smallnest/rpcx/server"
	"github.com/zylikedream/galaxy/core/gxyconfig"
	"github.com/zylikedream/galaxy/core/gxyregister"
)

type tcpConfig struct {
	Addr string `toml:"addr"`
}

type tcpTransport struct {
	conf *tcpConfig
}

func newTcpTransport(c *gxyconfig.Configuration) (*tcpTransport, error) {
	tran := &tcpTransport{}
	conf := &tcpConfig{}
	if err := c.UnmarshalKey(tran.Type(), conf); err != nil {
		return nil, err
	}
	tran.conf = conf
	return tran, nil
}

func (t *tcpTransport) Addr() string {
	return t.conf.Addr
}

func (t *tcpTransport) Network() string {
	return "tcp"
}

func (t *tcpTransport) Option() server.OptionFn {
	return emptyOptin
}

func (t *tcpTransport) Type() string {
	return TRANSPORT_TYPE_TCP
}

func (t *tcpTransport) Build(c *gxyconfig.Configuration, args ...interface{}) (interface{}, error) {
	return newTcpTransport(c)
}

func init() {
	gxyregister.Register((*tcpTransport)(nil))
}
