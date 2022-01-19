//  Copyright (c) 2018 Uber Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tario

import (
	"io"

	"github.com/klauspost/pgzip"
)


var _gzipCompressionLevelMap = map[string]int{
	"no":      pgzip.NoCompression,
	"speed":   pgzip.BestSpeed,
	"size":    pgzip.BestCompression,
	"default": pgzip.DefaultCompression,
}

type GzipCompressorFactory struct {
	compressionLevel int
}

func NewGzipCompressorFactory() *GzipCompressorFactory {
	return &GzipCompressorFactory {
		// Use gzip upstream default.
		compressionLevel: pgzip.DefaultCompression,
	}
}

// NewWriter returns a new gzip writer with gzip compression level.
func (c GzipCompressorFactory) NewWriter(w io.Writer) (io.WriteCloser, error) {
	return pgzip.NewWriterLevel(w, c.compressionLevel)
}

// NewReader returns a new gzip reader.
func (c GzipCompressorFactory) NewReader(r io.Reader) (io.ReadCloser, error) {
	return pgzip.NewReader(r)
}
