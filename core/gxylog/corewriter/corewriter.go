package corewriter

import (
	"io"

	"github.com/zylikedream/galaxy/core/gxyconfig"
	"github.com/zylikedream/galaxy/core/gxyregister"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	WRITER_TYPE_STDERR     = "writer.stderr"
	WRITER_TYPE_ROTATEFILE = "writer.rotate_file"
)

// WriterBuilder 根据key初始化writer
// Writer 日志interface
type CoreWriter interface {
	zapcore.Core
	io.Closer
}

type CloseFunc func() error

// Close 关闭
func (c CloseFunc) Close() error {
	return c()
}

var noopCloseFunc = func() error {
	return nil
}

func NewCoreWriter(t string, c *gxyconfig.Configuration, atomiclv zap.AtomicLevel) (CoreWriter, error) {
	if node, err := gxyregister.NewNode("writer."+t, c, atomiclv); err != nil {
		return nil, err
	} else {
		return node.(CoreWriter), nil
	}
}
