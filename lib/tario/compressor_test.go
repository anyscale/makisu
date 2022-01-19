package tario

import (
	"testing"

	"github.com/klauspost/pgzip"
	lz4 "github.com/pierrec/lz4/v4"
	"github.com/stretchr/testify/require"
)

func TestSetCompressionOptions(t *testing.T) {
	require := require.New(t)

	err := SetCompressionAlgorithm("invalid")
	require.Error(err)
	err = SetCompressionAlgorithm("gZiP")
	require.Error(err)

	err = SetCompressionAlgorithm("gzip")
	require.NoError(err)
	err = SetCompressionLevelGzip("invalid")
	require.Error(err)

	err = SetCompressionAlgorithm("gzip")
	require.NoError(err)
	err = SetCompressionLevelGzip("default")
	require.NoError(err)
	require.Equal(
		pgzip.DefaultCompression,
		algoToCompressorFactoryMap[compressionAlgo].(*GzipCompressorFactory).compressionLevel,
	)
	err = SetCompressionLevelGzip("size")
	require.NoError(err)
	require.Equal(
		pgzip.BestCompression,
		algoToCompressorFactoryMap[compressionAlgo].(*GzipCompressorFactory).compressionLevel,
	)

	err = SetCompressionAlgorithm("lz4")
	require.NoError(err)
	err = SetCompressionLevelLZ4("fast")
	require.NoError(err)
	require.Equal(
		lz4.Fast,
		algoToCompressorFactoryMap[compressionAlgo].(*LZ4CompressorFactory).compressionLevel,
	)
	err = SetCompressionLevelLZ4("level9")
	require.NoError(err)
	require.Equal(
		lz4.Level9,
		algoToCompressorFactoryMap[compressionAlgo].(*LZ4CompressorFactory).compressionLevel,
	)
}