// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by "gen.bash" from internal/trace; DO NOT EDIT.

//go:build go1.23

package trace

import "github.com/nathan-joslin/exp/trace/internal/version"

// GoVersion is the version set in the trace header.
func (r *Reader) GoVersion() version.Version {
	return r.version
}
