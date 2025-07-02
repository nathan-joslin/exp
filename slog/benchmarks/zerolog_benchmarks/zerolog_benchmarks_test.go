// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package zerolog_benchmarks

import (
	"io"
	"testing"

	slogbench "github.com/nathan-joslin/exp/slog/benchmarks"
	"github.com/rs/zerolog"
)

// Keep in sync (same names and behavior) as the
// benchmarks in the parent directory.

func BenchmarkAttrs(b *testing.B) {
	logger := zerolog.New(zerolog.SyncWriter(io.Discard)).With().Timestamp().Logger()
	b.Run("JSON discard", func(b *testing.B) {
		b.Run("5 args", func(b *testing.B) {
			b.ReportAllocs()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					logger.Info().
						Str("string", slogbench.TestString).
						Int("status", slogbench.TestInt).
						Dur("duration", slogbench.TestDuration).
						Time("time", slogbench.TestTime).
						Err(slogbench.TestError).
						Msg(slogbench.TestMessage)
				}
			})
		})
		b.Run("10 args", func(b *testing.B) {
			b.ReportAllocs()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					logger.Info().
						Str("string", slogbench.TestString).
						Int("status", slogbench.TestInt).
						Dur("duration", slogbench.TestDuration).
						Time("time", slogbench.TestTime).
						Err(slogbench.TestError).
						Str("string", slogbench.TestString).
						Int("status", slogbench.TestInt).
						Dur("duration", slogbench.TestDuration).
						Time("time", slogbench.TestTime).
						Err(slogbench.TestError).
						Msg(slogbench.TestMessage)
				}
			})
		})
	})
}
