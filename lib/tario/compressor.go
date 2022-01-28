package tario

import (
	"fmt"
	"io"

	"github.com/uber/makisu/lib/docker/image"
)

var compressionAlgoRead = "gzip"

var compressionAlgoWrite = "gzip"

var gzipCompressorFactory = NewGzipCompressorFactory()

var lz4CompressorFactory = NewLZ4CompressorFactory()

var zstdCompressorFactory = NewZstdCompressorFactory()

var algoToCompressorFactoryMap = map[string]CompressorFactory {
	"gzip": gzipCompressorFactory,
	"lz4": lz4CompressorFactory,
	"zstd": zstdCompressorFactory,
}

type CompressorFactory interface {
	NewWriter(w io.Writer) (io.WriteCloser, error)
	NewReader(r io.Reader) (io.ReadCloser, error)
}

// SetCompressionAlgorithmRead sets the compression algothm to be used for reading base images.
// Note: it doesn't check if the compression level being set is for the specific algorithm.
func SetCompressionAlgorithmRead(algo string) error {
	_, ok :=  algoToCompressorFactoryMap[algo]
	if !ok {
		return fmt.Errorf("invalid compression algorithm for read %s", algo)
	}
	compressionAlgoRead = algo

	return nil
}

// SetCompressionAlgorithmRead sets the compression algothm to be used for building current images.
// Note: it doesn't check if the compression level being set is for the specific algorithm.
func SetCompressionAlgorithmWrite(algo string) error {
	_, ok :=  algoToCompressorFactoryMap[algo]
	if !ok {
		return fmt.Errorf("invalid compression algorithm for read %s", algo)
	}
	compressionAlgoWrite = algo

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

// GetLayerMediaType returns docker layer media type based on the specified compression algo.
func GetLayerMediaType() string {
	switch compressionAlgoWrite {
	case "gzip":
		return image.MediaTypeLayerGzip
	case "lz4":
		// There is no lz4 support for media type. This is likely result in image to be rejected by docker.
		return image.MediaTypeLayerGzip
	case "zstd":
		return image.MediaTypeLayerZstd
	default:
		return ""
	}
}

// SetCompressionLevelZstd sets compression level for zstd compressor.
func SetCompressionLevelZstd(zstdCompressionLevelStr string) error {
	level, ok := _zstdCompressionLevelMap[zstdCompressionLevelStr]
	if !ok {
		return fmt.Errorf("invalid compression level for zstd: %s", zstdCompressionLevelStr)
	}
	zstdCompressorFactory.compressionLevel = level
	return nil
}

// NewCompressionReader returns a new compression reader with specified compression algorithm and
// compression level.
func NewCompressionReader(r io.Reader) (io.ReadCloser, error) {
	cf, ok :=  algoToCompressorFactoryMap[compressionAlgoRead]
	if !ok {
		return nil, fmt.Errorf("No suitable compressor factory for algorithm: %s", compressionAlgoRead)
	}

	return cf.NewReader(r)
}

// NewCompressionWriter returns a new compression writer with specified compression algorithm and
// compression level.
func NewCompressionWriter(w io.Writer) (io.WriteCloser, error) {
	cf, ok :=  algoToCompressorFactoryMap[compressionAlgoWrite]
	if !ok {
		return nil, fmt.Errorf("No suitable compressor factory for algorithm: %s", compressionAlgoWrite)
	}

	return cf.NewWriter(w)
}
