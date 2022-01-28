package tario

import (
	"fmt"
	"io"

	"github.com/klauspost/compress/zstd"
)

var _zstdCompressionLevelMap = map[string]zstd.EncoderLevel{
	"fastest":        zstd.SpeedFastest,
	"default":      zstd.SpeedDefault,
	"bettercompression":      zstd.SpeedBetterCompression,
	"bestcompression":      zstd.SpeedBestCompression,
}

type ZstdCompressorFactory struct {
	compressionLevel zstd.EncoderLevel
}

func NewZstdCompressorFactory() *ZstdCompressorFactory {
	return &ZstdCompressorFactory{
		// Use zstd upstream default.
		compressionLevel: zstd.SpeedDefault,
	}
}

// NewWriter returns a new zstd writer with specified zstd compression level.
func (c ZstdCompressorFactory) NewWriter(w io.Writer) (io.WriteCloser, error) {
	zstdW, err := zstd.NewWriter(w, zstd.WithEncoderLevel(c.compressionLevel), zstd.WithEncoderConcurrency(1))
	return zstdW, err
}

// NewReader returns a new zstd reader.
func (c ZstdCompressorFactory) NewReader(r io.Reader) (io.ReadCloser, error) {
	zstdR, err := zstd.NewReader(r)
	if err != nil {
		return nil, fmt.Errorf("init zstd reader: %s", err)
	}
	return io.NopCloser(zstdR), nil
}
