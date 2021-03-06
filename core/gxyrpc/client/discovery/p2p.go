package discovery

import (
	"github.com/smallnest/rpcx/client"
	"github.com/zylikedream/galaxy/core/gxyconfig"
	"github.com/zylikedream/galaxy/core/gxyregister"
)

type p2pDiscovery struct {
	conf *p2pConfig
	d    client.ServiceDiscovery
}

type p2pConfig struct {
	Peer string `toml:"peer"` // ex:tcp@localhost:2786
	Meta string `toml:"meta"`
}

func newP2pDiscovery(c *gxyconfig.Configuration) (*p2pDiscovery, error) {
	conf := &p2pConfig{}
	p2p := &p2pDiscovery{
		conf: conf,
	}
	if err := c.UnmarshalKey(p2p.Type(), conf); err != nil {
		return nil, err
	}
	d, err := client.NewPeer2PeerDiscovery(conf.Peer, conf.Meta)
	if err != nil {
		return nil, err
	}
	p2p.d = d
	return p2p, nil
}

func (p *p2pDiscovery) Type() string {
	return DISCOVERY_TYPE_P2P
}

func (p *p2pDiscovery) GetDiscovery() client.ServiceDiscovery {
	return p.d
}

func (p *p2pDiscovery) Build(c *gxyconfig.Configuration, args ...interface{}) (interface{}, error) {
	return newP2pDiscovery(c)
}

func init() {
	gxyregister.Register((*p2pDiscovery)(nil))
}
