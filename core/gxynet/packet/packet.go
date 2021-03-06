package packet

import (
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/zylikedream/galaxy/core/gxyconfig"
	"github.com/zylikedream/galaxy/core/gxynet/message"
	"github.com/zylikedream/galaxy/core/gxyregister"
)

var (
	ErrPkgHeadNotEnough = errors.New("pkg header not enougth")
	ErrPkgBodyNotEnough = errors.New("pkg body not enougth")
)

type PacketCodec interface {
	Decode(data []byte) (uint64, *message.Message, error) // 解析包体
	Encode(msg *message.Message) ([]byte, error)          // 打包消息
	Type() string
}

func convertUint(v uint64, len int) interface{} {
	switch len {
	case 1:
		return uint8(v)
	case 2:
		return uint16(v)
	case 4:
		return uint32(v)
	case 8:
		return v
	}
	return v
}

func uintDecode(data []byte, byteOrder binary.ByteOrder) (uint64, error) {
	switch len(data) {
	case 1:
		return uint64(data[0]), nil
	case 2:
		return uint64(byteOrder.Uint16(data)), nil
	case 4:
		return uint64(byteOrder.Uint32(data)), nil
	case 8:
		return uint64(byteOrder.Uint64(data)), nil
	}
	return 0, fmt.Errorf("unsupport byte len:%d", len(data))
}

const (
	PACKET_LTIV = "packet.ltiv"
)

func NewPacketCodec(t string, c *gxyconfig.Configuration) (PacketCodec, error) {
	if node, err := gxyregister.NewNode("packet."+t, c); err != nil {
		return nil, err
	} else {
		return node.(PacketCodec), nil
	}
}
