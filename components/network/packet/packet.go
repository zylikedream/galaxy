package packet

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/zylikedream/galaxy/components/gconfig"
	"github.com/zylikedream/galaxy/components/gregister"
	"github.com/zylikedream/galaxy/components/network/message"
)

type PacketCodec interface {
	ReadPacketLen(r io.Reader) (uint64, error)           // 读取消息长度
	DecodeBody(payLoad []byte) (*message.Message, error) // 解析包体
	Encode(msg *message.Message) ([]byte, error)         // 打包消息
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

func Uint(data []byte, byteOrder binary.ByteOrder) (uint64, error) {
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
	PACKET_LTIV = "ltiv"
)

var reg = gregister.NewRegister()

func Register(builder gregister.Builder) {
	reg.Register(builder)
}

func NewPacketCodec(t string, c *gconfig.Configuration) (PacketCodec, error) {
	if node, err := reg.NewNode(t, c); err != nil {
		return nil, err
	} else {
		return node.(PacketCodec), nil
	}
}
