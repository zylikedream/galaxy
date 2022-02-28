package processor

import (
	"github.com/zylikedream/galaxy/core/gxyconfig"
	"github.com/zylikedream/galaxy/core/gxynet/message"
	"github.com/zylikedream/galaxy/core/gxynet/packet"
	"github.com/zylikedream/galaxy/core/gxyregister"
)

type pathProcessor struct {
	pktCodec packet.PacketCodec
	msgCodec message.MessageCodec
	conf     *processorConfig
}

func newPathProcessor(c *gxyconfig.Configuration) (Processor, error) {
	proc := &pathProcessor{}
	conf := &processorConfig{}
	var err error
	if err = c.UnmarshalKey(proc.Type(), conf); err != nil {
		return nil, err
	}
	proc.conf = conf
	proc.pktCodec, err = packet.NewPacketCodec(conf.PacketCodecType, c)
	if err != nil {
		return nil, err
	}
	proc.msgCodec, err = message.NewMessageCodec(conf.MessageCodecType, c)
	if err != nil {
		return nil, err
	}
	return proc, nil
}

func (c *pathProcessor) Type() string {
	return "processor.path"
}

func (p *pathProcessor) Decode(data []byte) (uint64, *message.Message, error) {
	pkgLen, msg, err := p.pktCodec.Decode(data)
	if err == packet.ErrPkgBodyNotEnough || err == packet.ErrPkgHeadNotEnough { // 数据不足够，不算错误
		return 0, nil, nil
	}
	return pkgLen, msg, err
}

func (p *pathProcessor) Encode(msg *message.Message) ([]byte, error) {
	var err error
	msg.Payload, err = p.msgCodec.Encode(msg.Msg)
	if err != nil {
		return nil, err
	}
	data, err := p.pktCodec.Encode(msg)
	if err != nil {
		return nil, err
	}
	return data, err

}

func (p *pathProcessor) GetMessageCodec() message.MessageCodec {
	return p.msgCodec
}

func (p *pathProcessor) Build(c *gxyconfig.Configuration, args ...interface{}) (interface{}, error) {
	return newPathProcessor(c)
}

func init() {
	gxyregister.Register((*pathProcessor)(nil))
}
