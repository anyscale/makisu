package tario

import (
	"fmt"
	"io"

	lz4 "github.com/pierrec/lz4/v4"
)

var _lz4CompressionLevelMap = map[string]lz4.CompressionLevel{
	"fast":        lz4.Fast,
	"level1":      lz4.Level1,
	"level2":      lz4.Level2,
	"level3":      lz4.Level3,
	"level4":      lz4.Level4,
	"level5":      lz4.Level5,
	"level6":      lz4.Level6,
	"level7":      lz4.Level7,
	"level8":      lz4.Level8,
	"level9":      lz4.Level9,
}

type LZ4CompressorFactory struct {
	compressionLevel lz4.CompressionLevel
}

func NewLZ4CompressorFactory() *LZ4CompressorFactory {
	return &LZ4CompressorFactory{
		// Use lz4 upstream default.
		compressionLevel: lz4.Fast,
	}
}

// NewWriter returns a new lz4 writer with specified lz4 compression level.
func (c LZ4CompressorFactory) NewWriter(w io.Writer) (io.WriteCloser, error) {
	lz4W :=  lz4.NewWriter(w)
	if err := lz4W.Apply(lz4.CompressionLevelOption(c.compressionLevel)); err != nil {
		return nil, fmt.Errorf("Unable to initialize l4z writer with option %s: %s", c.compressionLevel, err)
	}
	return lz4W, nil
}

// NewReader returns a new lz4 reader.
func (c LZ4CompressorFactory) NewReader(r io.Reader) (io.ReadCloser, error) {
	return io.NopCloser(lz4.NewReader(r)), nil
}
