package tario

import (
	"testing"

	"github.com/klauspost/pgzip"
	lz4 "github.com/pierrec/lz4/v4"
	"github.com/stretchr/testify/require"
)

func TestSetCompressionOptions(t *testing.T) {
	require := require.New(t)

	err := SetCompressionAlgorithmRead("invalid")
	require.Error(err)
	err = SetCompressionAlgorithmWrite("gZiP")
	require.Error(err)

	err = SetCompressionAlgorithmRead("gzip")
	require.NoError(err)
	err = SetCompressionLevelGzip("invalid")
	require.Error(err)

	err = SetCompressionAlgorithmWrite("gzip")
	require.NoError(err)
	err = SetCompressionLevelGzip("default")
	require.NoError(err)
	require.Equal(
		pgzip.DefaultCompression,
		algoToCompressorFactoryMap[compressionAlgoWrite].(*GzipCompressorFactory).compressionLevel,
	)
	err = SetCompressionLevelGzip("size")
	require.NoError(err)
	require.Equal(
		pgzip.BestCompression,
		algoToCompressorFactoryMap[compressionAlgoWrite].(*GzipCompressorFactory).compressionLevel,
	)

	err = SetCompressionAlgorithmWrite("lz4")
	require.NoError(err)
	err = SetCompressionLevelLZ4("fast")
	require.NoError(err)
	require.Equal(
		lz4.Fast,
		algoToCompressorFactoryMap[compressionAlgoWrite].(*LZ4CompressorFactory).compressionLevel,
	)
	err = SetCompressionLevelLZ4("level9")
	require.NoError(err)
	require.Equal(
		lz4.Level9,
		algoToCompressorFactoryMap[compressionAlgoWrite].(*LZ4CompressorFactory).compressionLevel,
	)
}