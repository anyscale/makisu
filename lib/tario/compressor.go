package tario

import (
	"fmt"
	"io"
)

var compressionAlgo = "gzip"

var gzipCompressorFactory = NewGzipCompressorFactory()

var lz4CompressorFactory = NewLZ4CompressorFactory()

var algoToCompressorFactoryMap = map[string]CompressorFactory {
	"gzip": gzipCompressorFactory,
	"lz4": lz4CompressorFactory,
}

type CompressorFactory interface {
	NewWriter(w io.Writer) (io.WriteCloser, error)
	NewReader(r io.Reader) (io.ReadCloser, error)
}

// SetCompressionAlgorithm sets the compression algothm to be used for current build.
// Note: it doesn't check if the compression level being set is for the specific algorithm.
func SetCompressionAlgorithm(algo string) error {
	_, ok :=  algoToCompressorFactoryMap[algo]
	if !ok {
		return fmt.Errorf("invalid compression algorithm %s", algo)
	}
	compressionAlgo = algo

	return nil
}

// SetCompressionLevelGzip sets compression level for gzip compressor.
func SetCompressionLevelGzip(compressionLevelStr string) error {
	level, ok := _gzipCompressionLevelMap[compressionLevelStr]
	if !ok {
		return fmt.Errorf("invalid compression level for gzip: %s", compressionLevelStr)
	}
	gzipCompressorFactory.compressionLevel = level
	return nil
}

// SetCompressionLevelLZ4 sets compression level for lz4 compressor.
func SetCompressionLevelLZ4(lz4CompressionLevelStr string) error {
	level, ok := _lz4CompressionLevelMap[lz4CompressionLevelStr]
	if !ok {
		return fmt.Errorf("invalid compression level for lz4: %s", lz4CompressionLevelStr)
	}
	lz4CompressorFactory.compressionLevel = level
	return nil
}

// NewCompressionWriter returns a new compression writer with specified compression algorithm and
// compression level.
func NewCompressionWriter(w io.Writer) (io.WriteCloser, error) {
	cf, ok :=  algoToCompressorFactoryMap[compressionAlgo]
	if !ok {
		return nil, fmt.Errorf("No suitable compressor factory for algorithm: %s", compressionAlgo)
	}

	return cf.NewWriter(w)
}

// NewCompressionReader returns a new compression reader with specified compression algorithm and
// compression level.
func NewCompressionReader(r io.Reader) (io.ReadCloser, error) {
	cf, ok :=  algoToCompressorFactoryMap[compressionAlgo]
	if !ok {
		return nil, fmt.Errorf("No suitable compressor factory for algorithm: %s", compressionAlgo)
	}

	return cf.NewReader(r)
}
