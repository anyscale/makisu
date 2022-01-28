package tario

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestZstdWriteAndRead(t *testing.T) {
	require := require.New(t)

	testString := "testString"

	testFile, err := ioutil.TempFile("/tmp", "testGzipWriteAndRead")
	require.NoError(err)

	f := NewZstdCompressorFactory()
	w, err := f.NewWriter(testFile)
	require.NoError(err)
	w.Write([]byte(testString))
	w.Close()

	testFile.Seek(0, 0)

	r, err := f.NewReader(testFile)
	require.NoError(err)
	buffer := make([]byte, len([]byte(testString)))
	r.Read(buffer)

	require.Equal(testString, string(buffer))
}
